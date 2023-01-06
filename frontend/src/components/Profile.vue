<template>
    <div>
        <v-container fluid>
            <v-col fluid>
                <v-card class="mx-auto my-4" max-width="500">
                    <v-card-title class="text-h6 font-weight-regular justify-space-between">
                        <span>{{ currentTitle }}</span>
                        <v-avatar color="primary lighten-2" class="subheading white--text" size="24" v-text="step"></v-avatar>
                    </v-card-title>

                    <v-window v-model="step">
                        <v-window-item :value="1">
                            <v-card-text>
                                <v-text-field label="Email" value="davidjeremienoe@sjtu.edu.cn"></v-text-field>
                                <span class="text-caption grey--text text--darken-1"> This is the email you will use to login </span>
                            </v-card-text>
                        </v-window-item>

                        <v-window-item :value="2">
                            <v-card-text>
                                <v-text-field label="Password" type="password"></v-text-field>
                                <v-text-field label="Confirm Password" type="password"></v-text-field>
                                <span class="text-caption grey--text text--darken-1"> Please enter a password for your account </span>
                            </v-card-text>
                        </v-window-item>

                        <v-window-item :value="3">
                            <div class="pa-4 text-center">
                                <h3 class="text-h6 font-weight-light mb-2">Welcome to Älykäs koti</h3>
                                <span class="text-caption grey--text">Thanks for signing up!</span>
                            </div>
                        </v-window-item>
                    </v-window>

                    <v-card-actions>
                        <v-btn :disabled="step === 1" text @click="step--"> Clear </v-btn>
                        <v-spacer></v-spacer>
                        <v-btn :disabled="step === 3" color="primary" depressed @click="step++"> Next </v-btn>
                    </v-card-actions>
                </v-card>
                <v-spacer></v-spacer>
                <v-card class="mx-auto my-4" max-width="500">
                    <v-card-title class="text-h6 font-weight-regular justify-space-between">
                        <span>Sign In</span>
                    </v-card-title>

                    <v-window>
                        <v-window-item>
                            <v-card-text>
                                <v-text-field label="Email" value="davidjeremienoe@sjtu.edu.cn"></v-text-field>
                                <span class="text-caption grey--text text--darken-1"> This is the email you use to login </span>
                                <v-text-field label="Password" type="password"></v-text-field>
                                <span class="text-caption grey--text text--darken-1"> Password for your account </span>
                            </v-card-text>
                        </v-window-item>
                    </v-window>

                    <v-card-actions>
                        <v-btn text @click="resetForm"> Clear </v-btn>
                        <v-spacer></v-spacer>
                        <v-btn color="primary" depressed @click="login"> Login </v-btn>
                    </v-card-actions>
                </v-card>
            </v-col>
        </v-container>
    </div>
</template>

<script>
export default {
    name: "Profile",
    data: () => ({
        rules: {
            required: (value) => !!value || "Required.",
            counter: (value) => value.length <= 20 || "Max 20 characters",
            email: (value) => {
                const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
                return pattern.test(value) || "Invalid e-mail.";
            },
        },
        step: 1,
    }),
    methods: {
        resetForm() {},
        login() {},
        setGlobalTitle: function () {
            var myElement = document.getElementById("global-title");
            myElement.textContent = "Profile";
        },
    },
    mounted: function () {
        this.setGlobalTitle();
    },
    computed: {
        currentTitle() {
            switch (this.step) {
                case 1:
                    return "Sign-up";
                case 2:
                    return "Create a password";
                default:
                    return "Account created";
            }
        },
    },
};
</script>
