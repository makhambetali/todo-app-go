console.log("JS загружен!");

document.addEventListener("DOMContentLoaded", function () {
  console.log("DOM загружен!");
  loadTasks();

  document.getElementById("addTask").addEventListener("click", async function () {
    const text = document.getElementById("taskText").value.trim();
    const priority = document.getElementById("taskPriority").value;
    const dueDate = document.getElementById("taskDueDate").value;

    if (!text) {
      alert("Введите текст задачи!");
      return;
    }

    try {
      const response = await window.go.main.App.AddTask(text, priority, dueDate);
      console.log("Задача добавлена:", response);
      document.getElementById("taskText").value = "";
      loadTasks();
    } catch (error) {
      console.error("Ошибка при добавлении задачи:", error);
      alert("Ошибка при добавлении задачи!");
    }
  });
});


async function loadTasks() {
  try {
    const tasks = await window.go.main.App.GetTasks();
    const taskList = document.getElementById("taskList");
    const completedList = document.getElementById("completedTaskList");
    
    taskList.innerHTML = "";
    completedList.innerHTML = "";

    tasks.forEach((task) => {
      const li = document.createElement("li");
      li.className = task.done ? "done" : "";
      li.innerHTML = `
        <span class="task-text" onclick="toggleTask(${task.id})">
          ${task.text} <span class="priority">[${task.priority}]</span> - ${task.due_date}
        </span>
        <div class="task-buttons">
          <button class="toggle-btn" onclick="toggleTask(${task.id})">${task.done ? "✅" : "⬜"}</button>
          <button class="delete-btn" onclick="confirmDelete(${task.id})">🗑</button>
        </div>
      `;

      if (task.done) {
        completedList.appendChild(li);
      } else {
        taskList.appendChild(li);
      }
    });
  } catch (error) {
    console.error("Ошибка загрузки задач:", error);
  }
}

async function toggleTask(id) {
  try {
    await window.go.main.App.ToggleTask(id);
    loadTasks();
  } catch (error) {
    console.error("Ошибка изменения статуса задачи:", error);
  }
}

function confirmDelete(id) {
  const modal = document.getElementById("deleteModal");
  modal.style.display = "block";

  document.getElementById("confirmDelete").onclick = function () {
    deleteTask(id);
    modal.style.display = "none";
  };

  document.getElementById("cancelDelete").onclick = function () {
    modal.style.display = "none";
  };
}


async function deleteTask(id) {
  try {
    await window.go.main.App.DeleteTask(id);
    loadTasks();
  } catch (error) {
    console.error("Ошибка удаления задачи:", error);
  }
}


window.deleteTask = deleteTask;
window.toggleTask = toggleTask;
window.confirmDelete = confirmDelete;
