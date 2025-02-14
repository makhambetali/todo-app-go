# Функционал
1) Создание задач с выбором приоритета и сроком исполнения
2) Удаление задач с подтверждающим модальным окном
3) Отметка выполнения задач(To Do -> Done and Done - > To Do)
4) Валидация ввода
5) Сохранение данных после перезапуска программы: использование базы данных sqlite3

Command for wails installation: go install github.com/wailsapp/wails/v2/cmd/wails@latest

git clone https://github.com/makhambetali/todo-app-go.git
cd todo-app

# Установка фронтенда
cd frontend
npm install
npm run dev
cd ..

# Сборка Wails
wails dev


