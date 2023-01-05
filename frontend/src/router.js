import Vue from "vue";
import VueRouter from "vue-router";
import Dashboard from "./views/Dashboard.vue";
import Sensors from "./views/Sensors.vue";
import Settings from "./views/Settings.vue";
import History from "./views/History.vue";
import Automation from "./views/Automation.vue";
import Map from "./views/Map.vue";
import Debugger from "./views/Debugger.vue";
import Profile from "./views/Profile.vue"
Vue.use(VueRouter);

const routes = [
    {
        path: "/",
        name: "home",
        component: Dashboard,
    },
    {
        path: "/dashboard",
        name: "dashboard",
        component: Dashboard,
    },
    {
        path: "/sensors",
        name: "sensors",
        component: Sensors,
    },
    {
        path: "/settings",
        name: "settings",
        component: Settings,
    },
    {
        path: "/history",
        name: "history",
        component: History,
    },
    {
        path: "/automation",
        name: "automation",
        component: Automation,
    },
    {
        path: "/map",
        name: "map",
        component: Map,
    },
    {
        path: "/debugger",
        name: "debugger",
        component: Debugger,
    },
    {
        path: "/profile",
        name: "profile",
        component: Profile,
    },
];

const router = new VueRouter({
    mode: "history",
    routes,
});

export default router;
