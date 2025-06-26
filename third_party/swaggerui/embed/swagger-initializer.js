window.onload = function () {
    const configUrl = "/q/services"
    fetch(configUrl).then(r => r.text().then(t => {
        const urls = JSON.parse(t);

        window.ui = SwaggerUIBundle({
            urls: urls,
            dom_id: '#swagger-ui',
            deepLinking: true,
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIStandalonePreset
            ],
            plugins: [
                SwaggerUIBundle.plugins.DownloadUrl
            ],
            layout: "StandaloneLayout"
        });
    }));
};