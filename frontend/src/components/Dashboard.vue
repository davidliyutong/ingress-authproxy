<template>
  <div>
    <v-container>
      <v-card class="mx-auto my-4" max-width="400" dense>
        <v-list-item two-line>
          <v-list-item-content>
            <v-list-item-title class="text-h5" id="location"> Shanghai </v-list-item-title>
            <v-list-item-subtitle id="time-weather">Wen, Mostly sunny</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-card-text>
          <v-row align="center">
            <v-col cols="6" id="weather">
              <v-img src="https://cdn.vuetifyjs.com/images/cards/sun.png" alt="Sunny image" width="92"></v-img>
            </v-col>
            <v-col class="text-h3" cols="6" id="temperature"> 13&deg;C </v-col>
          </v-row>
        </v-card-text>

        <v-list-item>
          <v-list-item-icon>
            <v-icon>mdi-send</v-icon>
          </v-list-item-icon>
          <v-list-item-subtitle id="wind">3 km/h</v-list-item-subtitle>
        </v-list-item>

        <v-list-item>
          <v-list-item-icon>
            <v-icon>mdi-cloud</v-icon>
          </v-list-item-icon>
          <v-list-item-subtitle id="cloud">48%</v-list-item-subtitle>
        </v-list-item>
        <v-divider></v-divider>
      </v-card>
      <v-card class="mx-auto my-4" max-width="400" dense>
        <v-card-title class="text-h5"> Alarm </v-card-title>

        <v-card-text center>
          <v-col cols="6">
            <v-icon large>{{ alarmIcon }}</v-icon>
          </v-col>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn class="mx-2 my-2" color="success" depressed @click="alarm_state = 1">Enable</v-btn>
          <v-btn class="mx-2 my-2" color="error" depressed @click="alarm_state = 0">Disable</v-btn>
        </v-card-actions>
        <v-divider></v-divider>
      </v-card>
      <v-card class="mx-auto my-4" max-width="400" dense>
        <v-card-title class="text-h5"> Light </v-card-title>

        <v-card-text center>
          <v-list>
            <v-list-item>
              <v-row cols="6">
                Light 1
                <v-spacer></v-spacer>
                <v-switch v-model="switch1" @change="switchOnOff(switch1)"></v-switch>
              </v-row>
            </v-list-item>
            <v-list-item>
              <v-row cols="6">
                Light 2
                <v-spacer></v-spacer>
                <v-switch v-model="switch2"></v-switch>
              </v-row>
            </v-list-item>
          </v-list>
        </v-card-text>
        <v-divider></v-divider>
      </v-card>
      <v-card class="mx-auto my-4" max-width="400">
        <v-container center>
          <apexchart width="100%" type="bar" :options="chartOptions" :series="series" center></apexchart>
        </v-container>
      </v-card>
    </v-container>
  </div>
</template>

<script>
import axios from "axios";
import { mdiAlarmLight } from "@mdi/js";

export default {
  name: "Dashboard",
  data: () => ({
    switch1: false,
    switch2: false,
    alarm_state: 0,
    chartOptions: {
      chart: {
        id: "vuechart-example",
      },
      xaxis: {
        categories: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
      },
    },
    series: [
      {
        name: "series-1",
        data: [30, 40, 35, 50, 49, 60, 70],
      },
    ],
    console: console,
  }),

  mounted: function () {
    this.setGlobalTitle();
  },
  methods: {
    setGlobalTitle: function () {
      var myElement = document.getElementById("global-title");
      myElement.textContent = "Dashboard";
    },
    enableAlarm: function () {
      this.data.alarm_state = 1;
    },
    disableAlarm: function () {
      this.data.alarm_state = 0;
    },
    switchOnOff: function (s) {
      const Http = new XMLHttpRequest();
      if (s) {
        console.log("ON");
        Http.open("GET", "http://218.193.190.221:20889/on");
        Http.send();
      } else {
        console.log("OFF");
        Http.open("GET", "http://218.193.190.221:20889/off");
        Http.send();
      }
      Http.onreadystatechange = (e) => {
        console.log(Http.responseText);
      };
    },
  },
  computed: {
    alarmIcon() {
      console.log(this.alarm_state);
      switch (this.alarm_state) {
        case 0:
          return "mdi-alarm-light-off";
        case 1:
          return "mdi-alarm-light";
        default:
          return "mdi-alarm-light";
      }
    },
  },
};
</script>
