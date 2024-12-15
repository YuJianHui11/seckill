import { createStore } from 'vuex'
import { State } from './types'

export default createStore<State>({
  state: {
    user: null,
    token: null
  },
  mutations: {
    setUser(state: State, user: any) {
      state.user = user
    },
    setToken(state: State, token: string | null) {
      state.token = token
    }
  }
}) 