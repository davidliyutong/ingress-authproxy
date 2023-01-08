<template>
  <div>
  </div>
</template>

<script>
export default {
  name: "Secrets",
  data: () => ({}),
  mounted: function () {
    this.setGlobalTitle();
  },
  methods: {
    resetForm() {
    },
    login() {
    },
    setGlobalTitle: function () {
      var myElement = document.getElementById("global-title");
      myElement.textContent = "Secrets";
    },
  },
};
</script>
