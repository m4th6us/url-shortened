document.getElementById("urlForm").addEventListener("submit", async function(event) {
    event.preventDefault();

    const originalURL = document.getElementById("originalURL").value;
    const resultDiv = document.getElementById("result");

    try {
        // Realiza a requisição POST para o backend
        const response = await fetch("http://localhost:8080/shorten", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ original: originalURL })
        });

        if (!response.ok) {
            throw new Error("Erro ao encurtar URL");
        }

        const data = await response.json();

        // Exibe a URL encurtada
        resultDiv.innerHTML = `
            <p><strong>URL Encurtada:</strong></p>
            <a href="${data.short_url}" target="_blank">${data.short_url}</a>
        `;
    } catch (error) {
        resultDiv.innerHTML = `<p style="color: red;">Erro: ${error.message}</p>`;
    }
});
