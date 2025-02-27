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

# Default view
![image](https://github.com/user-attachments/assets/1ac59448-b390-41c6-ae3b-3977be490505)

# Input Validation
![image](https://github.com/user-attachments/assets/b78f2e4f-cf9e-4c83-906b-eaef2dfcb257)

# Delete Confirmation
![image](https://github.com/user-attachments/assets/291ad892-bf62-4382-933b-822b5c87804d)

# Done tasks
![image](https://github.com/user-attachments/assets/0e935c5e-40d8-4ef0-9bac-25eac3ca7891)




