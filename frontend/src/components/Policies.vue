<template>
    <div>

    </div>
</template>

<script>
export default {
    name: "Policies",
    data: () => ({}),
    mounted: function () {
        this.setGlobalTitle();
    },
    methods: {
        setGlobalTitle: function () {
            var myElement = document.getElementById("global-title");
            myElement.textContent = "Policies";
        },
    },
};
</script>
