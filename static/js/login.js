// document.getElementById("submit").addEventListener("click", function(event) {
//     event.preventDefault(); // Formun normal gönderimini engelleyin

//     var username = document.getElementById("username").value;
//     var password = document.getElementById("password").value;
    
//     if (username && password) {
//         alert("Username: " + username + "\nPassword: " + password);
//         // postMessage fonksiyonunuzu burada çağırın
//         // postMessage("/login", {username: username, password: password});
//     }
// });

// document.getElementById("search").addEventListener("click", function() {
//     var search = document.getElementById("search-input").value;
//     if (search) {
//         window.location.href = "/search/" + search;
//     }
// });

document.getElementById("submit").addEventListener("click", function(event) {
    event.preventDefault(); // Formun normal gönderimini engelleyin

    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    
    if (username && password) {
        // Sunucuya fetch ile POST isteği gönder
        fetch("/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({username: username, password: password})
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            alert(data.message);
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }
});

document.getElementById("search").addEventListener("click", function() {
    var search = document.getElementById("search-input").value;
    if (search) {
        window.location.href = "/search/" + search;
    }
});
