<template>
    <v-app>
        <v-navigation-drawer v-model="drawer" app bottom>
            <v-list-item>
                <v-list-item-content>
                    <v-list-item-title class="text-h6">AuthProxy</v-list-item-title>
                    <v-list-item-subtitle>Yet another ingress auth proxy</v-list-item-subtitle>
                </v-list-item-content>
            </v-list-item>
            <v-divider></v-divider>
            <v-list dense nav>
                <v-list-item link to="/dashboard">
                    <v-list-item-action>
                        <v-icon>mdi-view-dashboard</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Dashboard</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                <v-list-item link to="/sensors">
                    <v-list-item-action>
                        <v-icon>mdi-leak </v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Sensors</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                <v-list-item link to="/history">
                    <v-list-item-action>
                        <v-icon>mdi-history</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>History</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                <v-list-item link to="/automation">
                    <v-list-item-action>
                        <v-icon>mdi-auto-fix</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Automation</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                <v-list-item link to="/map">
                    <v-list-item-action>
                        <v-icon>mdi-map</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Map</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                <v-list-item link to="/debugger">
                    <v-list-item-action>
                        <v-icon>mdi-bug</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Debugger</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-list>
            <template v-slot:append>
                <v-list dense nav>
                    <v-list-item link to="/profile">
                        <v-list-item-action>
                            <v-icon>mdi-account</v-icon>
                        </v-list-item-action>
                        <v-list-item-content>
                            <v-list-item-title>Profile</v-list-item-title>
                        </v-list-item-content>
                    </v-list-item>
                    <v-list-item link to="/settings">
                        <v-list-item-action>
                            <v-icon>mdi-cog</v-icon>
                        </v-list-item-action>
                        <v-list-item-content>
                            <v-list-item-title>Settings</v-list-item-title>
                        </v-list-item-content>
                    </v-list-item>
                </v-list>
                <!-- <div class="pa-2">
                    <v-btn block>Logout</v-btn>
                </div> -->
            </template>
        </v-navigation-drawer>

        <v-app-bar app color="#6363A8" dark>
            <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
            <v-toolbar-title id="global-title"></v-toolbar-title>
            <v-spacer></v-spacer>

            <v-btn icon>
                <v-icon>mdi-magnify</v-icon>
            </v-btn>
            <v-btn icon link to="/dashboard">
                <v-icon>mdi-home</v-icon>
            </v-btn>
            <v-btn icon link to="/settings">
                <v-icon>mdi-dots-vertical</v-icon>
            </v-btn>
        </v-app-bar>

        <!-- 根据应用组件来调整你的内容 -->
        <v-main>
            <!-- 给应用提供合适的间距 -->
            <v-container fluid>
                <!-- 如果使用 vue-router -->
                <router-view></router-view>
            </v-container>
        </v-main>

        <v-footer color="#6363A8" app>
            <span class="white--text">&copy; 2021</span>
        </v-footer>
    </v-app>
</template>

<script>
export default {
    name: "App",
    data: function () {
        return {
            drawer: null,
        };
    },
};
</script>
