/** @type {import('tailwindcss').Config} */
export default {
    content: ["./templates/**/*.templ"],
    theme: {
        extend: {
            fontFamily: {
                mono: ["Courier Prime", "monospace"],
            },
        },
    },
    plugins: [],
    corePlugins: {
        preflight: true,
    },
};
