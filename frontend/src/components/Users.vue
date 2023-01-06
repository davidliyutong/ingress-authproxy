<template>
    <div>
        <v-expansion-panels focusable multiple>
            <v-expansion-panel v-for="(item, i) in 5" :key="i">
                <v-expansion-panel-header>Users</v-expansion-panel-header>
                <v-expansion-panel-content> Users Description </v-expansion-panel-content>
            </v-expansion-panel>
        </v-expansion-panels>
    </div>
</template>

<script>
export default {
    name: "Users",
    data: () => ({}),
    mounted: function () {
        this.setGlobalTitle();
    },
    methods: {
        setGlobalTitle: function () {
            var myElement = document.getElementById("global-title");
            myElement.textContent = "Users";
        },
    },
};
</script>
