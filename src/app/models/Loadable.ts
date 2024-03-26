export interface Loadable<T> {
  isLoading: boolean
  data: T
  error?: unknown
}
