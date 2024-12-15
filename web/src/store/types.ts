export interface State {
  user: any | null
  token: string | null
}

export interface RootState {
  state: State
} 