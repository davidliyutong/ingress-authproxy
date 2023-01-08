<template>
  <div>
    <v-container>
      <v-card
          class="mx-auto my-4"
      >
        <div class="pa-4 text-center">
          <h3 class="text-h6 font-weight-light mb-2">Welcome, {{ username }}</h3>
          <span class="text-caption grey--text">Have a great day !! </span>
        </div>
      </v-card>

      <v-spacer></v-spacer>

      <v-card class="mx-auto  my-4" >
        <v-list-item three-line>
          <v-list-item-content>
            <v-list-item-title class="headline mb-1">Version</v-list-item-title>
            <v-list-item-subtitle>{{ version }}</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-card-actions>
          <v-btn text @click="getVersion()">Refresh</v-btn>
        </v-card-actions>
      </v-card>

      <v-card class="mx-auto my-4" >
        <v-row align="center">
          <v-col>
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-title class="headline mb-1">Health</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-card-text class="display-3" cols="6">
              {{ healthOK }}
            </v-card-text>
          </v-col>
          <v-col>
            <v-list-item three-line>
              <v-list-item-content>
                <v-list-item-title class="headline mb-1">Response Time</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-card-text class="display-3" cols="6" color="healthColor">
              {{ healthTimeout }}
            </v-card-text>
          </v-col>
        </v-row>
        <v-card-actions>
          <v-btn text @click="getHealth()">Refresh</v-btn>
        </v-card-actions>
      </v-card>

    </v-container>
  </div>
</template>

<script>
import axios from "axios";
import {mdiAlarmLight} from "@mdi/js";

function getRootPath() {
  return window.location.protocol + '//' + window.location.host;
}

export default {
  name: "Dashboard",
  data: () => ({
    username: localStorage.getItem("username"),
    version: "",
    healthOK: false,
    healthTimeout: "-1ms",
    healthColor: "red"
  }),

  mounted: function () {
    this.setGlobalTitle();
  },
  methods: {
    initialize() {
      this.getVersion()
      this.getHealth()
    },
    getVersion() {
      let targetURL = getRootPath() + "/v1/version";
      axios.get(targetURL).then(response => {
        if (response.status === 200) {
          this.version = response.data
        }
      })
    },
    getHealth() {
      let targetURL = getRootPath() + "/v1/healthz";
      let now = Date.now();
      axios.get(targetURL).then(response => {
        let elapsed = Date.now() - now;
        if (response.status === 200) {
          this.healthOK = true
          if (elapsed <= 50) {
            this.healthColor = "green"
          } else {
            this.healthColor = "yellow"
          }
          this.healthTimeout = elapsed + "ms"
        } else {
          this.healthOK = false
          this.healthTimeout = "-1ms"
        }
      })
    },
    setGlobalTitle: function () {
      var myElement = document.getElementById("global-title");
      myElement.textContent = "Dashboard";
    },

  },
  computed: {},
  created() {
    this.initialize()
  },
};
</script>
