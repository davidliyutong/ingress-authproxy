<template>
    <div class="views-settings">
        <v-card> </v-card>
    </div>
</template>

<script>
export default {
    name: "Settings",
    data: () => ({}),
    mounted: function () {
        this.setGlobalTitle();
    },
    methods: {
        setGlobalTitle : function (){
            var myElement = document.getElementById("global-title");
            myElement.textContent = "Settings";
        }
    }
};
</script>
