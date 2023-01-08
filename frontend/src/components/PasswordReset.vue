<template>
  <v-app id="login">
    <v-main>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="primary">
                <v-toolbar-title>Password Reset</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <v-form ref="form" @submit.prevent="changePassword()">
                  <v-text-field
                      v-model="username"
                      name="username"
                      label="Username"
                      type="text"
                      placeholder="username"
                      required
                  ></v-text-field>

                  <v-text-field
                      v-model="oldpassword"
                      name="oldpassword"
                      label="Old Password"
                      type="password"
                      placeholder="oldpassword"
                      required
                  ></v-text-field>

                  <v-text-field
                      v-model="password"
                      name="password"
                      label="Password"
                      type="password"
                      placeholder="password"
                      :rules="rules"
                      required
                  ></v-text-field>

                  <v-text-field
                      v-model="confirm"
                      name="confirm"
                      label="Confirm"
                      type="password"
                      placeholder="confirm"
                      :rules="rules"
                      required
                  ></v-text-field>
                  <v-card-actions>
                    <v-btn class="mt-4" @click="resetForm">Clear</v-btn>
                    <v-spacer></v-spacer>
                    <v-btn type="submit" class="mt-4" color="primary" value="log in">Apply</v-btn>
                  </v-card-actions>
                  <v-snackbar
                      v-model="snackbarFail"
                      color="red"
                  >
                    Change Password Failed
                    <template v-slot:action="{ attrs }">
                      <v-btn
                          text
                          v-bind="attrs"
                          @click="snackbarFail = false"
                      >
                        Close
                      </v-btn>
                    </template>
                  </v-snackbar>
                  <v-snackbar
                      v-model="snackbarOK"
                      color="green"
                  >
                    Change Password Succeed
                    <template v-slot:action="{ attrs }">
                      <v-btn
                          text
                          v-bind="attrs"
                          @click="snackbarOK = false"
                      >
                        Close
                      </v-btn>
                    </template>
                  </v-snackbar>
                </v-form>
              </v-card-text>
            </v-card>

            <v-card
                class="mx-auto my-8 elevation-12"
                link to="/"
                color="success"
            >
              <v-container>
                <v-row>
                  <v-icon color="white" class="mx-4">mdi-arrow-left</v-icon>
                  <v-spacer></v-spacer>
                  <v-card-title class="mx-4">
                    <div class="white--text">
                      Back to Login
                    </div>
                  </v-card-title>
                </v-row>
              </v-container>
            </v-card>

          </v-flex>
        </v-layout>
      </v-container>
    </v-main>
  </v-app>

</template>

<script>
import qs from 'qs'
import axios from 'axios'

function getRootPath() {
  return window.location.protocol + '//' + window.location.host;
}

function b64EncodeUnicode(str) {
  return btoa(encodeURIComponent(str).replace(/%([0-9A-F]{2})/g,
      function (match, p1) {
        return String.fromCharCode('0x' + p1);
      }));
}

async function getUser(username, basicToken) {
  let targetURL = getRootPath() + "/v1/user/" + username;
  let user = null
  try {
    await axios.get(targetURL, {
      headers: {
        Authorization: basicToken,
      }
    }).then(response => {
      if (response.status === 200) {
        user = response.data
      }
    })
  } catch (err) {
    user = null
  }

  return user
}

async function updateUser(user, basicToken) {
  if (user === null) {
    return false
  }
  let targetURL = getRootPath() + "/v1/user/" + user.metadata.name;
  let succeed = false;
  try {
    await axios.put(targetURL, JSON.stringify(user),{
          headers: {
            Authorization: basicToken,
          }
        }
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
  name: "PasswordReset",
  data: function () {
    return {
      username: "",
      oldpassword: "",
      password: "",
      confirm: "",
      drawer: false,
      snackbarFail: "",
      snackbarOK: "",
    }
  },
  methods: {
    validateField: function () {
      this.$refs.form.validate()
    },
    resetForm: function () {
      this.$refs.form.reset()
    },
    changePassword: async function () {
      if (this.$refs.form.validate() === false){
        this.$message.bottom().error('New Password Mismatch')
        return
      }
      let basicToken = "Basic " + b64EncodeUnicode(this.username + ":" + this.oldpassword)

      let user = await getUser(this.username,  basicToken)
      if (user != null) {
        user.password = this.password
        let succeed = await updateUser(user, basicToken)
        if (succeed) {
          this.$message.bottom().success('Password Changed')
        }
      } else {
        this.$message.bottom().error('Password Change Failed')
      }
    }
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
