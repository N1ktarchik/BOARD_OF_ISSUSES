package handlers

import (
	"log"
	"net/http"
)

func getUserIDFromContext(r *http.Request) int {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Panicf("middleware context error")
	}

	return userID
}

func (h *UserHandler) HandleBase(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Board of Issues API</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 900px;
            margin: 0 auto;
            padding: 20px;
            background: #f5f5f5;
        }
        h1 {
            color: #2c3e50;
            border-bottom: 3px solid #3498db;
            padding-bottom: 10px;
        }
        h2 {
            color: #34495e;
            margin-top: 30px;
        }
        .endpoint {
            background: white;
            border-radius: 8px;
            padding: 15px;
            margin: 10px 0;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        .method {
            display: inline-block;
            padding: 5px 10px;
            border-radius: 4px;
            font-weight: bold;
            margin-right: 10px;
        }
        .get { background: #61affe; color: white; }
        .post { background: #49cc90; color: white; }
        .patch { background: #fca130; color: white; }
        .delete { background: #f93e3e; color: white; }
        .path {
            font-family: monospace;
            font-size: 1.1em;
            color: #2c3e50;
        }
        .description {
            margin-top: 10px;
            color: #666;
        }
        code {
            background: #f8f9fa;
            padding: 2px 5px;
            border-radius: 3px;
            font-family: monospace;
        }
        .auth-note {
            background: #fff3cd;
            border-left: 4px solid #ffc107;
            padding: 15px;
            margin: 20px 0;
        }
    </style>
</head>
<body>
    <h1>📋 Board of Issues API</h1>
    <p>Добро пожаловать в API сервиса управления задачами и досками!</p>
    
    <div class="auth-note">
        <strong>🔐 Аутентификация:</strong> Большинство эндпоинтов требуют Bearer токен в заголовке 
        <code>Authorization: Bearer &lt;ваш_токен&gt;</code>
    </div>

    <h2>👤 Пользователи</h2>
    
    <div class="endpoint">
        <span class="method post">POST</span>
        <span class="path">/register</span>
        <div class="description">Регистрация нового пользователя</div>
        <code>{"name": "Имя", "email": "user@mail.com", "password": "123456"}</code>
    </div>

    <div class="endpoint">
        <span class="method post">POST</span>
        <span class="path">/login</span>
        <div class="description">Вход в систему</div>
        <code>{"email": "user@mail.com", "password": "123456"}</code>
    </div>

    <div class="endpoint">
        <span class="method patch">PATCH</span>
        <span class="path">/api/users/name</span>
        <div class="description">Изменение имени пользователя</div>
        <code>{"name": "Новое имя"}</code>
    </div>

    <div class="endpoint">
        <span class="method patch">PATCH</span>
        <span class="path">/api/users/password</span>
        <div class="description">Изменение пароля</div>
        <code>{"old_password": "старый", "new_password": "новый"}</code>
    </div>

    <div class="endpoint">
        <span class="method patch">PATCH</span>
        <span class="path">/api/users/email</span>
        <div class="description">Изменение email</div>
        <code>{"email": "new@mail.com"}</code>
    </div>

    <div class="endpoint">
        <span class="method post">POST</span>
        <span class="path">/api/users</span>
        <div class="description">Подключение пользователя к доске</div>
        <code>{"desk_id": 123}</code>
    </div>

    <h2>📊 Доски</h2>

    <div class="endpoint">
        <span class="method post">POST</span>
        <span class="path">/api/desks</span>
        <div class="description">Создание новой доски</div>
        <code>{"name": "Моя доска", "password": "123456"}</code>
    </div>

    <div class="endpoint">
        <span class="method get">GET</span>
        <span class="path">/api/desks</span>
        <div class="description">Получение списка всех досок пользователя</div>
    </div>

    <div class="endpoint">
        <span class="method patch">PATCH</span>
        <span class="path">/api/desks/{id}/name</span>
        <div class="description">Изменение названия доски</div>
        <code>{"name": "Новое название"}</code>
    </div>

    <div class="endpoint">
        <span class="method patch">PATCH</span>
        <span class="path">/api/desks/{id}/password</span>
        <div class="description">Изменение пароля доски</div>
        <code>{"password": "новый_пароль"}</code>
    </div>

    <div class="endpoint">
        <span class="method patch">PATCH</span>
        <span class="path">/api/desks/{id}/owner</span>
        <div class="description">Смена владельца доски</div>
        <code>{"new_owner_id": 456}</code>
    </div>

    <div class="endpoint">
        <span class="method delete">DELETE</span>
        <span class="path">/api/desks/{id}</span>
        <div class="description">Удаление доски</div>
    </div>

    <h2>✅ Задачи</h2>

    <div class="endpoint">
        <span class="method post">POST</span>
        <span class="path">/api/tasks</span>
        <div class="description">Создание задачи</div>
        <code>{"title": "Задача", "description": "Описание", "desk_id": 123}</code>
    </div>

    <div class="endpoint">
        <span class="method get">GET</span>
        <span class="path">/api/tasks</span>
        <div class="description">Получение всех задач</div>
    </div>

    <div class="endpoint">
        <span class="method get">GET</span>
        <span class="path">/api/tasks?done=true&desk_id=123</span>
        <div class="description">Получение задач с фильтрацией</div>
    </div>

    <div class="endpoint">
        <span class="method patch">PATCH</span>
        <span class="path">/api/tasks/{id}/complyte</span>
        <div class="description">Отметить задачу как выполненную</div>
    </div>

    <div class="endpoint">
        <span class="method patch">PATCH</span>
        <span class="path">/api/tasks/{id}/time</span>
        <div class="description">Добавить время к задаче</div>
        <code>{"hours": 2}</code>
    </div>

    <div class="endpoint">
        <span class="method patch">PATCH</span>
        <span class="path">/api/tasks/{id}/description</span>
        <div class="description">Изменить описание задачи</div>
        <code>{"description": "Новое описание"}</code>
    </div>

    <div class="endpoint">
        <span class="method delete">DELETE</span>
        <span class="path">/api/tasks/{id}</span>
        <div class="description">Удаление задачи</div>
    </div>

    <footer style="margin-top: 50px; text-align: center; color: #999; font-size: 0.9em;">
        <p>Board of Issues API v1.0</p>
    </footer>
</body>
</html>
`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}
