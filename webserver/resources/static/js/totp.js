document.addEventListener('DOMContentLoaded', function () {
    populateTotpDetails()
}, false);

async function populateTotpDetails() {
    const response = await fetch("/register_mfa/totp/", {
        method: 'GET',
        mode: 'same-origin',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'follow'
    });

    if (response.ok) {

        let details;
        try {
            details = await response.json();
        } catch (e) {
            document.getElementById("serverError").hidden = false;
            return
        }

        document.getElementById("ImageData").src = details.ImageData;

        document.getElementById("AccountName").textContent = details.AccountName;
        document.getElementById("Key").textContent = details.Key;

    } else {
        document.getElementById("serverError").hidden = false;

        console.log("Unable to fetch TOTP details for registration: ", response.status, response.text);
    }
}