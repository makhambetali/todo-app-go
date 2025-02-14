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
