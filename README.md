git clone https://github.com/makhambetali/todo-app-go.git
cd todo-app

# Установка фронтенда
cd frontend
npm install
npm run dev
cd ..

# Сборка Wails
wails dev
