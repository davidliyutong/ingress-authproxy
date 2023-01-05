<template>
    <v-app>
        <v-navigation-drawer v-model="drawer" app>
            <v-list dense>
                <v-list-item link>
                    <v-list-item-action>
                        <v-icon>mdi-home</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Home</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                <v-list-item link to="/sensors">
                    <v-list-item-action>
                        <v-icon>mdi-motion-sensor </v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Sensors</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                <v-list-item link>
                    <v-list-item-action>
                        <v-icon>mdi-email</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Contact</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-list>
        </v-navigation-drawer>

        <v-app-bar app color="#007BFF" dark>
            <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
            <v-toolbar-title>IMU WebUI</v-toolbar-title>
        </v-app-bar>

        <!-- 根据应用组件来调整你的内容 -->
        <v-main>
            <!-- 给应用提供合适的间距 -->
            <v-container fluid>
                <!-- 如果使用 vue-router -->
                <router-view></router-view>
            </v-container>
        </v-main>

        <v-footer color="#007BFF" app>
            <span class="white--text">&copy; 2021</span>
        </v-footer>
    </v-app>
</template>

<script>
export default {
    name: "Home",
    data: () => ({
        drawer: null,
    }),
};
</script>
