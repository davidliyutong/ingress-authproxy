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
                <v-list-item link>
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

        <v-app-bar app color=#007BFF dark>
            <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
            <v-toolbar-title>IMU WebUI</v-toolbar-title>
        </v-app-bar>

        <v-main>
            <v-container class="fill-height" fluid>
                <v-row align="center" justify="center">
                    <v-col class="text-center">
                        <v-tooltip left>
                            <template v-slot:activator="{ on }">
                                <v-btn :href="source" icon large target="_blank" v-on="on">
                                    <v-icon small>mdi-code-tags</v-icon>
                                </v-btn>
                            </template>
                            <span>Source</span>
                        </v-tooltip>
                    </v-col>
                </v-row>
            </v-container>
        </v-main>
        <v-footer color="indigo" app>
            <span class="white--text">&copy; 2021</span>
        </v-footer>
    </v-app>
</template>

<script>
export default {
    name: 'Dashboard',
    data: () => ({
        drawer: null,
    }),
};
</script>
