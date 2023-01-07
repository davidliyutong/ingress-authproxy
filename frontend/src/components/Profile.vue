<template>
  <div>
    <v-container fluid>
      <v-tabs>
        <v-tab>Info</v-tab>
        <v-tab>Security</v-tab>
        <v-tab-item key="1">

        </v-tab-item>
        <v-tab-item key="2">
          <v-container>
            <v-card class="mx-auto my-4">
              <v-card-title class="text-h6 font-weight-regular justify-space-between">
                <span>Change Password</span>
              </v-card-title>

              <v-card-text>
                <v-form ref="form" @submit.prevent="changePassword()">
                  <v-text-field
                      v-model="password"
                      name="username"
                      label="New Password"
                      type="password"
                      placeholder="password"
                      :rules="rules"
                      required
                  ></v-text-field>

                  <v-text-field
                      v-model="confirm"
                      name="confirm"
                      label="Confirm Password"
                      type="password"
                      placeholder="confirm"
                      required
                  ></v-text-field>
                  <v-card-actions>
                    <v-btn class="mt-4" @click="resetForm">Clear</v-btn>
                    <v-spacer></v-spacer>
                    <v-btn type="submit" class="mt-4" color="primary" value="log in">Change Password</v-btn>
                  </v-card-actions>
                </v-form>
              </v-card-text>
            </v-card>
          </v-container>

        </v-tab-item>
      </v-tabs>
    </v-container>
  </div>
</template>

<script>

import axios from "axios";

function getRootPath() {
  return window.location.protocol + '//' + window.location.host;
}

async function getUser(username) {
  let targetURL = getRootPath() + "/v1/admin/users/" + username;
  let user = null
  try {
    await axios.get(targetURL,
    ).then(response => {
      if (response.status === 200) {
        user = response.data
      }
    })
  } catch (err) {
    user = null
  }

  return user
}

async function updateUser(user) {
  if (user === null) {
    return false
  }
  let targetURL = getRootPath() + "/v1/admin/users/" + user.name;
  let succeed = false;
  try {
    await axios.put(targetURL, JSON.stringify(user)
    ).then(response => {
      if (response.status === 200) {
        succeed = true
      }
    })
  } catch (err) {
    succeed = false
  }

  return succeed
}

export default {
  name: "Profile",
  data: () => ({
    password: "",
    confirm: "",
  }),
  methods: {
    resetForm: function () {
      this.$refs.form.reset()
    },
    changePassword: async function () {
      if (this.$refs.form.validate() === false){

        this.$message.bottom().error('New Password Mismatch')
        return
      }
      let user = await getUser(localStorage.getItem('username'))
      if (user != null) {
        user.password = this.password
        // console.log(user)
        let succeed = await updateUser(user)
        if (succeed) {
          this.$message.bottom().success('Password Changed')
        }
      } else {
        this.$message.bottom().error('Password Change Failed')
      }
    },
    setGlobalTitle: function () {
      var myElement = document.getElementById("global-title");
      myElement.textContent = "Profile";
    },
    validateField: function () {
      this.$refs.form.validate()
    }
    ,
  },
  mounted: function () {
    this.setGlobalTitle();
  },
  computed: {
    rules() {
      const rules = []
      rules.push(v => (v || '').length <= 32 ||
          'A maximum of 32 characters is allowed')

      rules.push(v => (v || '').indexOf(' ') < 0 ||
          'No spaces are allowed')

      if (this.confirm) {
        rules.push(v => (!!v && v) === this.confirm ||
            'Values do not match')
      }

      if (this.password) {
        rules.push(v => (!!v && v) === this.password ||
            'Values do not match')
      }

      return rules
    },
  },

  watch: {
    confirm: 'validateField',
    password: 'validateField',
  },
};
</script>
