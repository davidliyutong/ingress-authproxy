<template>
  <div>
    <v-container>
      <v-window>
      <v-window-item :value="5">
        <div class="pa-4 text-center">
          <h3 class="text-h6 font-weight-light mb-2">Welcome, {{ username }}</h3>
          <span class="text-caption grey--text">Have a great day !! </span>
        </div>
      </v-window-item>
      </v-window>
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
    username: localStorage.getItem("username")
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
