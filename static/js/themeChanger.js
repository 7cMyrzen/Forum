function loadTheme() {
    // Verifier si theme existe dans localstrorage sinon le creer
    if (localStorage.getItem('theme') === null) {
        localStorage.setItem('theme', 'normal');
    }

    //Charger le theme normal
    if (localStorage.getItem('theme') === 'normal') {
        // Changer les couleurs des variables root
        document.documentElement.style.setProperty('--color-primary', '#ffffff');
        document.documentElement.style.setProperty('--color-secondary', '#000000');
        document.documentElement.style.setProperty('--color-tertiary', '#d9d9d9');
        document.documentElement.style.setProperty('--align-items-icon', 'flex-start');
        document.documentElement.style.setProperty('--normal-icon', 'block');
        document.documentElement.style.setProperty('--dark-icon', 'none');
        document.documentElement.style.setProperty('--theme-bg', '#d9d9d9');
        document.documentElement.style.setProperty('--box-shadow', '0px 4px 4px 0px rgba(0, 0, 0, 0.25) inset;');

        //Changer le theme dans le localstorage
        localStorage.setItem('theme', 'normal');
    }

    //Charger le theme dark

    if (localStorage.getItem('theme') === 'dark') {
        // Changer les couleurs des variables root
        document.documentElement.style.setProperty('--color-primary', '#202020');
        document.documentElement.style.setProperty('--color-secondary', '#d50162');
        document.documentElement.style.setProperty('--color-tertiary', '#d9d9d9');
        document.documentElement.style.setProperty('--align-items-icon', 'flex-end');
        document.documentElement.style.setProperty('--normal-icon', 'none');
        document.documentElement.style.setProperty('--dark-icon', 'block');
        document.documentElement.style.setProperty('--theme-bg', '#202020');

        localStorage.setItem('theme', 'dark');
    }
}

loadTheme();

function changeTheme() {
    if (localStorage.getItem('theme') === 'normal') {
        localStorage.setItem('theme', 'dark');
    } else {
        localStorage.setItem('theme', 'normal');
    }

    loadTheme();
}