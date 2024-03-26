import { Loadable } from '@/app/models/Loadable'

export interface KubeconfigLocation {
  name: string
  location: string
  connected: boolean
}

export interface RawConfig {
  content: string
}

export interface FileConfig {
  location: string
}

export interface KubeconfigState {
  kubeconfigLocations: Loadable<Array<KubeconfigLocation>>
  selectedLocation?: string
}
