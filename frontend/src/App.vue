<template>
  <v-container fluid>
  <router-view></router-view>
  </v-container>
</template>

<script>
import axios from "axios";
import { mdiAlarmLight } from "@mdi/js";

export default {
  name: "Admin",
  data: function () {
    return {
      drawer: null,
    };
  },
  methods: {
  },
  computed: {
  },
};
</script>
