module.exports = {
    content: [
        "./src/**/*.{js,jsx,ts,tsx}",
        "./public/index.html",
    ],
    theme: {
        extend: {},
    },
    plugins: [require("daisyui")], // Include daisyUI if you are using it
};