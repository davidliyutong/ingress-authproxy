<template>
    <div>
        <v-expansion-panels focusable multiple>
            <v-expansion-panel v-for="(item, i) in 5" :key="i">
                <v-expansion-panel-header>Sensor</v-expansion-panel-header>
                <v-expansion-panel-content> Sensor Description </v-expansion-panel-content>
            </v-expansion-panel>
        </v-expansion-panels>
    </div>
</template>

<script>
export default {
    name: "Sensors",
    data: () => ({}),
    mounted: function () {
        this.setGlobalTitle();
    },
    methods: {
        setGlobalTitle: function () {
            var myElement = document.getElementById("global-title");
            myElement.textContent = "Sensors";
        },
    },
};
</script>
