<template>
  <v-tabs>
    <v-tab>Info</v-tab>
    <v-tab>Operations</v-tab>
    <v-tab-item key="1">
      <v-container>
        <v-card>
          <v-card-title>
            Server Options
          </v-card-title>
          <v-card-text>

            <p>
              {{ option }}
            </p>

          </v-card-text>
          <v-card-actions>
            <v-btn text @click="getOption()">Refresh</v-btn>
          </v-card-actions>
        </v-card>


      </v-container>


    </v-tab-item>
    <v-tab-item key="2">
      <v-container>


        <v-card class="mx-auto  my-4">
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="headline mb-1">Sync Policies</v-list-item-title>
              <v-list-item-subtitle>Sync policies from backend database</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
          <v-card-actions>
            <v-btn color="success" text @click="sync()">Sync</v-btn>
          </v-card-actions>
        </v-card>


        <v-card class="mx-auto  my-4">
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="headline mb-1">Shutdown</v-list-item-title>
              <v-list-item-subtitle>Shutdown the backend server, will reboot this instance under docker deployment</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
          <v-card-actions>
            <v-btn color="error" text @click="shutdown()">Shutdown</v-btn>
          </v-card-actions>
        </v-card>

      </v-container>

    </v-tab-item>
  </v-tabs>
</template>

<script>
import axios from "axios";

function getRootPath() {
  return window.location.protocol + '//' + window.location.host;
}

export default {
  name: "Settings",
  data: () => ({
    option: "null"
  }),
  mounted: function () {
    this.setGlobalTitle();
  },
  methods: {
    setGlobalTitle: function () {
      var myElement = document.getElementById("global-title");
      myElement.textContent = "Settings";
    },
    shutdown() {
      let targetURL = getRootPath() + "/v1/admin/server/shutdown"

      try {
        axios.post(targetURL, "").then(
            response => {
              if (response.status === 200) {
                this.$message.success("Execute Server Shutdown")
              } else {
                this.$message.error("Failed to Execute Server Shutdown")
              }
            }
        )

      } catch (err) {
        this.$message.error("Failed to Execute Server Shutdown")

      }

    },
    sync() {
      let targetURL = getRootPath() + "/v1/admin/server/sync"

      try {
        axios.post(targetURL, "").then(
            response => {
              if (response.status === 200) {
                this.$message.success("Sync Triggered")
              } else {
                this.$message.error("Failed to Sync")
              }
            }
        )

      } catch (err) {
        this.$message.error("Failed to Sync")

      }

    },
    getOption() {
      let targetURL = getRootPath() + "/v1/admin/server/option"
      try {
        axios.get(targetURL).then(response => {
          if (response.status === 200) {
            this.option = JSON.stringify(response.data, null, "    ")
          } else {
            this.option = "null"
          }

        })
      } catch (err) {
        // console.log("err")
        this.option = "null"
      }
    }
  },
  created() {
    this.getOption()
  }
};
</script>
