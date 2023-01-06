<template>
    <div>
        <v-card class="mx-auto my-4">
            <v-timeline>
                <v-timeline-item color="red lighten-2">Last Person Left Home</v-timeline-item>
                <v-timeline-item class="text-right" color="#6363A8">Turn Off Lights</v-timeline-item>
            </v-timeline>
        </v-card>
        <v-card class="mx-auto my-4">
            <v-timeline>
                <v-timeline-item color="red lighten-2">Temperature Drops Below 25</v-timeline-item>
                <v-timeline-item class="text-right" color="#6363A8">Turn on AC</v-timeline-item>
                <v-timeline-item color="#6363A8">Close Windows</v-timeline-item>
            </v-timeline>
        </v-card>
    </div>
</template>

<script>
export default {
    name: "Automation",
    data: () => ({
        el: 1,
    }),
    mounted: function () {
        this.setGlobalTitle();
    },
    methods: {
        setGlobalTitle: function () {
            var myElement = document.getElementById("global-title");
            myElement.textContent = "Automation";
        },
    },
};
</script>
