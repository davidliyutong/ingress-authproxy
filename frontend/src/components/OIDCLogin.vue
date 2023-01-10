<template>
  <v-app id="login">
    <v-main>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="orange">
                <v-icon>mdi-openid</v-icon>
                <v-spacer></v-spacer>
                <v-toolbar-title>OIDC Login</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
               Auth Request ID: {{ id }}
              </v-card-text>
              <v-card-text class="align-center">
                <form method="POST" action="/oidc/login/username">
                  <input
                      type="hidden" name="id" :value="id">
                  <div class="v-input theme--light v-text-field v-text-field--is-booted v-text-field--placeholder">
                    <div class="v-input__control">
                      <div class="v-input__slot">
                        <div class="v-text-field__slot"><label for="input-52" class="v-label theme--light"
                                                               style="left: 0px; right: auto; position: absolute;">Username</label><input
                            name="username" required="required" id="input-52" type="text"></div>
                      </div>
                      <div class="v-text-field__details">
                        <div class="v-messages theme--light">
                          <div class="v-messages__wrapper"></div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="v-input theme--light v-text-field v-text-field--is-booted v-text-field--placeholder">
                    <div class="v-input__control">
                      <div class="v-input__slot">
                        <div class="v-text-field__slot"><label for="input-55" class="v-label theme--light"
                                                               style="left: 0px; right: auto; position: absolute;">Password</label><input
                            name="password" required="required" id="input-55" type="password"></div>
                      </div>
                      <div class="v-text-field__details">
                        <div class="v-messages theme--light">
                          <div class="v-messages__wrapper"></div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="v-card__actions">
                    <button type="button"
                            class="mt-4 v-btn v-btn--is-elevated v-btn--has-bg theme--light v-size--default"><span
                        class="v-btn__content">Clear</span></button>
                    <div class="spacer"></div>
                    <button type="submit" value="log in"
                            class="mt-4 v-btn v-btn--is-elevated v-btn--has-bg theme--light v-size--default orange">
                      <span class="v-btn__content">Login</span></button>
                  </div>
                  <div class="v-snack v-snack--bottom v-snack--has-background">
                    <div class="v-snack__wrapper v-sheet theme--dark red" style="display: none;">
                      <div role="status" aria-live="polite" class="v-snack__content"> Login Failed</div>
                      <div class="v-snack__action ">
                        <button type="button" class="v-btn v-btn--text theme--dark v-size--default v-snack__btn"><span
                            class="v-btn__content"> Close </span></button>
                      </div>
                    </div>
                  </div>
                  <div class="v-snack v-snack--bottom v-snack--has-background">
                    <div class="v-snack__wrapper v-sheet theme--dark green" style="display: none;">
                      <div role="status" aria-live="polite" class="v-snack__content"> Login Succeed</div>
                      <div class="v-snack__action ">
                        <button type="button" class="v-btn v-btn--text theme--dark v-size--default v-snack__btn"><span
                            class="v-btn__content"> Close </span></button>
                      </div>
                    </div>
                  </div>
                </form>
<!--                <form method="POST" action="/oidc/login/username" style="height: 200px; width: 200px;">-->
<!--                  <div>-->
<!--                    <label for="username">Username:</label>-->
<!--                    <input id="username" name="username" style="width: 100%">-->
<!--                  </div>-->

<!--                  <div>-->
<!--                    <label for="password">Password:</label>-->

<!--                    <input id="password" name="password" style="width: 100%">-->
<!--                  </div>-->
<!--                  <button type="submit">Login</button>-->
<!--                </form>-->
                <!--                <form ref="form" @submit.prevent="login()">-->
                <!--                  <v-text-field-->
                <!--                      v-model="username"-->
                <!--                      name="username"-->
                <!--                      label="Username"-->
                <!--                      type="text"-->
                <!--                      placeholder="username"-->
                <!--                      required-->
                <!--                  ></v-text-field>-->

                <!--                  <v-text-field-->
                <!--                      v-model="password"-->
                <!--                      name="password"-->
                <!--                      label="Password"-->
                <!--                      type="password"-->
                <!--                      placeholder="password"-->
                <!--                      required-->
                <!--                  ></v-text-field>-->
                <!--                  <v-card-actions>-->
                <!--                    <v-btn class="mt-4" @click="resetForm">Clear</v-btn>-->
                <!--                    <v-spacer></v-spacer>-->
                <!--                    <v-btn type="submit" class="mt-4" color="orange" value="log in">Login</v-btn>-->
                <!--                  </v-card-actions>-->
                <!--                  <v-snackbar-->
                <!--                      v-model="snackbarFail"-->
                <!--                      color="red"-->
                <!--                  >-->
                <!--                    Login Failed-->
                <!--                    <template v-slot:action="{ attrs }">-->
                <!--                      <v-btn-->
                <!--                          text-->
                <!--                          v-bind="attrs"-->
                <!--                          @click="snackbarFail = false"-->
                <!--                      >-->
                <!--                        Close-->
                <!--                      </v-btn>-->
                <!--                    </template>-->
                <!--                  </v-snackbar>-->
                <!--                  <v-snackbar-->
                <!--                      v-model="snackbarOK"-->
                <!--                      color="green"-->
                <!--                  >-->
                <!--                    Login Succeed-->
                <!--                    <template v-slot:action="{ attrs }">-->
                <!--                      <v-btn-->
                <!--                          text-->
                <!--                          v-bind="attrs"-->
                <!--                          @click="snackbarOK = false"-->
                <!--                      >-->
                <!--                        Close-->
                <!--                      </v-btn>-->
                <!--                    </template>-->
                <!--                  </v-snackbar>-->
                <!--                </form>-->
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

function getRootPath() {
  return window.location.protocol + '//' + window.location.host;
}


function GetQueryString(name) {
  let reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
  let r = window.location.search.substring(1).match(reg);
  if (r != null) return decodeURI(r[2]);
  return null;
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
      id: GetQueryString("authRequestID")
    }
  },
  methods: {
    resetForm: function () {
      this.$refs.form.reset()
    },
    login: async function () {
      let data = {
        id: GetQueryString("authRequestID"),
        username: this.username,
        password: this.password,
      }
      console.log(data)
      let options = {
        config: {headers: {'Content-Type': 'application/x-www-form-urlencoded'}},
        data: qs.stringify(data),
        url: getRootPath() + '/oidc/login/username',
      };
      console.log(options)
      axios.post(options.url, options.data, options.config
      ).then(response => {
        console.log(response)
        if (response.status === 200) {
          if (response.data.code === 200) {
            console.log("Login Succeed")
            this.snackbarOK = true
            this.$router.push(response.data.callback_url)
            return
          }
        }
        this.snackbarFail = true
      })
    }
  },
};
</script>
