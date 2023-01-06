import Vue from "vue";
import Vuetify from "vuetify";
import "vuetify/dist/vuetify.min.css";

Vue.use(Vuetify);

Vue.component("image-icon", {
    template: `<v-img :src="imgsrc" max-width="32"></v-img>`,
    props: ["imgsrc"],
});

export default new Vuetify({
    theme: {
        themes: {
            light: {
                primary: "#3265A8",
                secondary: "#424242",
                accent: "#82B1FF",
                error: "#FF5252",
                info: "#2196F3",
                success: "#4CAF50",
                warning: "#FFC107",
            },
        },
    },
    icons: {
        iconfont: "mdi", // 设置使用本地的icon资源
        values: {
            sensor: {
                component: "image-icon",
                props: {
                    imgsrc: "https://vuetifyjs.com/favicon.ico",
                },
            },
        },
    },
});
