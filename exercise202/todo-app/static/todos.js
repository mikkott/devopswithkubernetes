fetch('http://localhost:8081/todos')
    .then(response => response.json())
    .then(data => {
        const todos = data;
        const list = document.createElement("ul");
        todos.forEach(todo => {
            const listItem = document.createElement("li");
            listItem.textContent = todo.title;
            list.appendChild(listItem);
        });
        document.body.appendChild(list);
    })
    .catch(error => {
        console.error('Error fetching todos:', error);
    });