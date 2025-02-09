document.addEventListener("DOMContentLoaded", function () {
    console.log("script.js cargado correctamente"); // Para verificar que el script se ejecuta

    // Validación del formulario
    const form = document.getElementById("contactForm");
    const errorMessage = document.getElementById("error-message");

    if (form) {
        form.addEventListener("submit", function (event) {
            let errors = [];

            const name = document.getElementById("name").value.trim();
            const email = document.getElementById("email").value.trim();
            const age = document.getElementById("age").value;

            // Validar nombre
            if (name.length < 3) {
                errors.push("El nombre debe tener al menos 3 caracteres.");
            }

            // Validar email con expresión regular
            const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
            if (!emailRegex.test(email)) {
                errors.push("El correo electrónico no es válido.");
            }

            // Validar edad (entre 18 y 100)
            if (age < 18 || age > 100) {
                errors.push("La edad debe estar entre 18 y 100 años.");
            }

            // Mostrar errores si existen
            if (errors.length > 0) {
                event.preventDefault(); // Evitar el envío del formulario
                errorMessage.innerHTML = errors.join("<br>");
            } else {
                errorMessage.innerHTML = "";
            }
        });
    }

    // Manejo del menú desplegable
    const menuToggle = document.getElementById("menu-toggle");
    const menu = document.getElementById("menu");

    if (menuToggle && menu) {
        menuToggle.addEventListener("click", function () {
            menu.classList.toggle("active");
            console.log("Menú desplegado:", menu.classList.contains("active")); // Verifica si el menú se activa
        });
    } else {
        console.log("Error: No se encontró el botón o el menú.");
    }
});
