<template>
  <v-app id="login">
    <v-main>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="primary">
                <v-toolbar-title>Admin Login</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <form ref="form" @submit.prevent="login()">
                  <v-text-field
                      v-model="username"
                      name="username"
                      label="Username"
                      type="text"
                      placeholder="username"
                      required
                  ></v-text-field>

                  <v-text-field
                      v-model="password"
                      name="password"
                      label="Password"
                      type="password"
                      placeholder="password"
                      required
                  ></v-text-field>
                  <v-card-actions>
                    <v-btn class="mt-4" @click="resetForm">Clear</v-btn>
                    <v-spacer></v-spacer>
                    <v-btn type="submit" class="mt-4" color="primary" value="log in">Login</v-btn>
                  </v-card-actions>
                  <v-snackbar
                      v-model="snackbarFail"
                      color="red"
                  >
                    Login Failed
                    <template v-slot:action="{ attrs }">
                      <v-btn
                          text
                          v-bind="attrs"
                          @click="snackbar = false"
                      >
                        Close
                      </v-btn>
                    </template>
                  </v-snackbar>
                  <v-snackbar
                      v-model="snackbarOK"
                      color="green"
                  >
                    Login Succeed
                    <template v-slot:action="{ attrs }">
                      <v-btn
                          text
                          v-bind="attrs"
                          @click="snackbar = false"
                      >
                        Close
                      </v-btn>
                    </template>
                  </v-snackbar>
                </form>
              </v-card-text>
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

async function getToken (username, password) {
  let params = {
    accessKey: username,
    secretKey: password
  };
  // console.log("username: ", params.accessKey, "password: ", params.secretKey)

  let paramStr = qs.stringify(params);
  let loginUrl = window.document.location.href + "v1/jwt/login";
  let loginSucceed = false
  try {
    await axios.post(loginUrl, paramStr
    ).then(response => {
      let responseResult = JSON.stringify(response.data)
      if (response.data.code === 200) {
        localStorage.setItem('accessKey', username);
        localStorage.setItem('token', responseResult.token)
        loginSucceed = true
      }
    })
  } catch(err) {
    loginSucceed = false
  }

  return loginSucceed
}

export default {
  name: "Login",
  data: function () {
    return {
      username: "",
      password: "",
      drawer: false,
      snackbarFail: false,
      snackbarOK: false,
    }
  },
  methods: {
    resetForm: function () {
      this.$refs.form.reset()
    },
    login: async function () {
      let loginSucceed = await getToken(this.username,this.password)

      console.log(loginSucceed)
      if (loginSucceed !== true) {
          this.snackbarFail = true
      } else {
        this.snackbarOK = true
        await this.$router.push('/admin')
      }

    }
  },
};
</script>
