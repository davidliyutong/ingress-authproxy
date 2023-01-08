<template>
  <div>
    <v-card class="mx-auto my-4">
      <v-card-title class="text-h6 font-weight-regular justify-space-between">
        <span>{{ currentTitle }}</span>
        <v-avatar color="primary lighten-2" class="subheading white--text" size="24" v-text="step"></v-avatar>
      </v-card-title>

      <v-window v-model="step">
        <v-window-item :value="1">
          <v-card-text>
            <v-text-field v-model="username" label="Username" value=""></v-text-field>
            <span class="text-caption grey--text text--darken-1"> This is the username you will use to login </span>
          </v-card-text>
        </v-window-item>

        <v-window-item :value="2">
          <v-card-text>
            <v-text-field v-model="password" label="Password" placeholder="password" type="password"></v-text-field>
            <v-text-field v-model="confirm" label="Confirm Password" placeholder="confirm" type="password"></v-text-field>
            <span class="text-caption grey--text text--darken-1"> Please enter a password for your account </span>
          </v-card-text>
        </v-window-item>

        <v-window-item :value="3">
          <v-card-text>

            <v-text-field v-model="phone" label="Phone" value=""></v-text-field>
            <v-text-field v-model="email" label="Email" value=""></v-text-field>
            <v-text-field v-model="nickname" label="Nickname" value=""></v-text-field>
            <v-checkbox
                v-model="isAdmin"
                label="Is Admin"
            ></v-checkbox>
          </v-card-text>
        </v-window-item>

        <v-window-item :value="4">
          <div class="pa-4 text-center">
            <h3 class="text-h6 font-weight-light mb-2">Confirm ?</h3>
            <span class="text-caption grey--text">The user will be added</span>
          </div>
        </v-window-item>

        <v-window-item :value="5">
          <div class="pa-4 text-center">
            <h3 class="text-h6 font-weight-light mb-2">User Added</h3>
            <span class="text-caption grey--text">The user is added</span>
          </div>
        </v-window-item>
      </v-window>

      <v-card-actions>
        <v-btn :disabled="step === 1" text @click="step--"> {{ userCreationLeftButtonText }}</v-btn>
        <v-spacer></v-spacer>
        <v-btn color="primary" depressed @click=createUserRightButton()> {{
            userCreationRightButtonText
          }}
        </v-btn>
      </v-card-actions>
    </v-card>


    <!--    <v-card-title>-->
    <!--      Users-->
    <!--      <v-spacer></v-spacer>-->
    <!--      <v-text-field-->
    <!--          v-model="search"-->
    <!--          append-icon="mdi-magnify"-->
    <!--          label="Search"-->
    <!--          single-line-->
    <!--          hide-details-->
    <!--      ></v-text-field>-->
    <!--    </v-card-title>-->
    <!--        <v-data-table-->
    <!--            :headers="headers"-->
    <!--            :items="userItems"-->
    <!--            :items-per-page="10"-->
    <!--            class="elevation-1"-->
    <!--            no-data-text="No Users"-->
    <!--            show-select-->
    <!--            item-key="name"-->
    <!--            :search="search"-->
    <!--        ></v-data-table>-->

    <v-data-table
        :headers="headers"
        :items="userItems"
        sort-by="calories"
        :items-per-page=20
        class="elevation-1"
        no-data-text="No Users"
        show-select
        :search="search"
    >
      <template v-slot:top>
        <v-toolbar flat color="white">
          <v-toolbar-title>Users</v-toolbar-title>
          <v-divider
              class="mx-4"
              inset
              vertical
          ></v-divider>
          <v-text-field
              v-model="search"
              append-icon="mdi-magnify"
              label="Search"
              single-line
              hide-details
          ></v-text-field>
          <v-spacer></v-spacer>
          <v-btn color="primary" @click="initialize">Refresh</v-btn>
          <v-dialog v-model="dialog" max-width="500px">
            <!--            <template v-slot:activator="{ on, attrs }">-->
            <!--              <v-btn-->
            <!--                  color="primary"-->
            <!--                  dark-->
            <!--                  class="mb-2"-->
            <!--                  v-bind="attrs"-->
            <!--                  v-on="on"-->
            <!--              >New User</v-btn>-->
            <!--            </template>-->
            <v-card>
              <v-card-title>
                <span class="headline">{{ formTitle }}</span>
              </v-card-title>

              <v-card-text>
                <v-container>
                  <v-text-field v-model="editedItem.nickname" label="Nickname"></v-text-field>
                  <v-text-field v-model="editedItem.phone" label="Phone"></v-text-field>
                  <v-text-field v-model="editedItem.email" label="Email"></v-text-field>
                  <v-text-field v-model="editedItem.status" label="Status"></v-text-field>
                  <v-text-field v-model="password" label="New Password" placeholder="password"
                                type="password"></v-text-field>
                  <v-text-field v-model="confirm" label="Confirm Password" placeholder="confirm"
                                type="confirm"></v-text-field>
                  <v-checkbox
                      v-model="editedItem.isAdmin"
                      label="Is Admin"
                  ></v-checkbox>
                </v-container>
              </v-card-text>

              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
                <v-btn color="blue darken-1" text @click="save">Save</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-toolbar>
      </template>
      <template v-slot:item.actions="{ item }">
        <v-icon
            small
            class="mr-2"
            @click="editItem(item)"
        >
          mdi-pencil
        </v-icon>
        <v-icon
            small
            @click="deleteItem(item)"
        >
          mdi-delete
        </v-icon>
      </template>
      <template v-slot:no-data>
        <v-btn color="primary" @click="initialize">Refresh</v-btn>
      </template>
    </v-data-table>
  </div>
</template>

<script>
import axios from "axios";

function getRootPath() {
  return window.location.protocol + '//' + window.location.host;
}

async function createUser(user) {
  if (user.nickname === "") {
    user.nickname = user.metadata.name
  }
  if (user.password !== user.confirm) {
    return false
  }
  if (user.isAdmin) {
    user.isAdmin = 1
  } else {
    user.isAdmin = 0
  }
  let targetURL = getRootPath() + "/v1/admin/users"
  let succeed = false;
  try {
    await axios.post(targetURL, JSON.stringify(user)).then(
        response => {
          succeed = response.status === 200;
        }
    )
  } catch (err) {
    succeed = false
  }

  return succeed
}

async function updateUser(user) {
  if (user === null) {
    return false
  }
  let targetURL = getRootPath() + "/v1/admin/users/" + user.metadata.name;
  let succeed = false;
  try {
    user.status = Number(user.status)
    if (user.isAdmin) {
      user.isAdmin = 1
    } else {
      user.isAdmin = 0
    }
    // console.log(user)
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

async function deleteUser(username) {
  if (username === "") {
    return false
  }
  let targetURL = getRootPath() + "/v1/admin/users/" + username;
  let succeed = false;
  try {
    await axios.delete(targetURL
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
  name: "Users",
  data: () => ({
    search: "",
    step: 1,
    username: "",
    password: "",
    confirm: "",
    phone: "",
    email: "",
    nickname: "",
    isAdmin: false,
    headers: [
      {
        text: 'name',
        align: 'start',
        value: 'metadata.name',
      },
      {text: 'nickname',value: 'nickname',},
      {text: 'Phone', value: 'phone'},
      {text: 'Email', value: 'email'},
      {text: 'Last Login', value: 'loginedAt'},
      {text: 'Status', value: 'status'},
      {text: 'Is Admin', value: 'isAdmin'},
      {text: 'Actions', value: 'actions', sortable: false},
    ],
    userItems: [],
    editedIndex: -1,
    editedItem: {
      name: '',
      nickname: '',
      phone: '',
      email: '',
      status: 1,
      isAdmin: 0,
    },
    defaultItem: {
      name: '',
      nickname: '',
      phone: '',
      email: '',
      status: 1,
      isAdmin: 0,
    },
    dialog: false,
  }),
  mounted: function () {
    this.setGlobalTitle();
  },
  methods: {
    initialize() {
      let targetURL = getRootPath() + "/v1/admin/users";
      try {
        axios.get(targetURL,
        ).then(response => {
          // console.log(JSON.stringify(response.data.items, null, 4))
          if (response.status === 200) {
            this.userItems = response.data.items
          } else {
            this.userItems = []
          }
        })
      } catch (err) {
        this.userItems = []
      }
    },
    editItem(item) {
      this.editedIndex = this.userItems.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
      // console.log(this.editedIndex, this.editedItem)
    },

    async deleteItem(item) {
      const index = this.userItems.indexOf(item)
      let ret = confirm('Are you sure you want to delete this item?')
      if (ret === true) {
        let succeed = await deleteUser(this.userItems[index].name)
        if (succeed) {
          this.$message.bottom().success("User Deleted")
          this.userItems.splice(index, 1)
        } else {
          this.$message.bottom().success("Failed to Delete User")

        }
      }
    },

    close() {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    async save() {
      if (this.editedIndex > -1) {
        Object.assign(this.userItems[this.editedIndex], this.editedItem)
      } else {
        this.userItems.push(this.editedItem)
      }
      let succeed = await updateUser(this.editedItem)
      if (succeed) {
        this.$message.bottom().success("User Updated")
      } else {
        this.$message.bottom().success("Failed to Update User")

      }
      this.close()
    },
    setGlobalTitle: function () {
      let myElement = document.getElementById("global-title");
      myElement.textContent = "Users";
    },
    async createUserRightButton() {
      if (this.step <= 3) {
        this.step++
        return null
      } else if (this.step === 4) {
        let user = {
          nickname: this.nickname,
          password: this.password,
          confirm: this.confirm,
          email: this.email,
          phone: this.phone,
          isAdmin: this.isAdmin,
          metadata: {
            name: this.username
          },
        }
        let succeed = await createUser(user)
        if (succeed) {
          this.step++
        } else {
          this.$message.bottom().error("Failed to Create User")
        }
      } else {
         this.step = 1
      }
    }
  },
  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'New User' : 'Edit User'
    },
    currentTitle() {
      switch (this.step) {
        case 1:
          return "Create User";
        case 2:
          return "Password";
        case 3:
          return "Other Information";
        default:
          return "";
      }
    },
    userCreationRightButtonText: function () {
      if (this.step === 3) {
        return "Create"
      } else if (this.step === 4) {
        return "Confirm"
      } else if (this.step === 5) {
        return "New"
      } else {
        return "Next"
      }
    },
    userCreationLeftButtonText: function () {
      if (this.step === 3) {
        return "Cancel"
      } else {
        return "Previous"
      }
    }
  },
  watch: {
    dialog(val) {
      val || this.close()
    },
  },
  created() {
    this.initialize()
  },

};
</script>
