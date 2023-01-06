<template>
  <div>
    <v-expansion-panels focusable multiple>
      <v-expansion-panel v-for="(item, i) in 5" :key="i">
        <v-expansion-panel-header>Secrets</v-expansion-panel-header>
        <v-expansion-panel-content> Secrets Description</v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
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
