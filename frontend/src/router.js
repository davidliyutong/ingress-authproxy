import Vue from "vue";
import VueRouter from "vue-router";
import Dashboard from "./components/Dashboard.vue"
import Users from "./components/Users.vue"
import Policies from "./components/Policies.vue"
import Secrets from "./components/Secrets.vue"

import Login from "./components/Login.vue";
import PasswordReset from "@/components/PasswordReset.vue";
import Profile from "@/components/Profile.vue";
import Settings from "@/components/Settings.vue";
import App from "@/App.vue";
import Admin from "@/components/Admin.vue";
import OIDCLogin from "@/components/OIDCLogin.vue";

Vue.use(VueRouter);

const original = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
    return original.call(this, location).catch(err => err)
}

const routes = [
    {
        path: "/",
        name: "login",
        component: Login,
    },
    {
        path: "/oidc-login",
        name: "oidc-login",
        component: OIDCLogin,
    },
    {
        path: "/passwordreset",
        name: "passwordreset",
        component: PasswordReset,
    },
    {
        path: "/admin",
        name: "admin",
        redirect: "/admin/dashboard",
        component: Admin,
        meta: {
            requiresAuth: true,
        },
        children: [
            {
                path: "/admin/dashboard",
                name: "dashboard",
                component: Dashboard,
                meta: {
                    requiresAuth: true,
                }
            },
            {
                path: "/admin/users",
                name: "users",
                component: Users,
                meta: {
                    requiresAuth: true,
                }
            },
            {
                path: "/admin/policies",
                name: "policies",
                component: Policies,
                meta: {
                    requiresAuth: true,
                }
            },
            {
                path: "/admin/secrets",
                name: "secrets",
                component: Secrets,
                meta: {
                    requiresAuth: true,
                }
            },
            {
                path: "/admin/profile",
                name: "profile",
                component: Profile,
                meta: {
                    requiresAuth: true,
                }
            },
            {
                path: "/admin/settings",
                name: "settings",
                component: Settings,
                meta: {
                    requiresAuth: true,
                }
            }
        ]
    },

];

const router = new VueRouter({
    mode: "history",
    routes,
    base: '/'
});

router.beforeEach((to, from, next) => {
    let token = localStorage.getItem("token")
    if (token == null || token === '') {
        if (to.path === '/' || to.path === '/passwordreset' || to.path === '/oidc-login') {
            next();
        } else {
            next({name: 'login'});
        }
    } else {
        if (to.path === '/') {
            next({name: 'admin'});
        } else {
            next();
        }
    }
})

export default router;
