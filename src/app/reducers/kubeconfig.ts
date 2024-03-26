import { createReducer, PayloadAction } from '@reduxjs/toolkit'
import { KubeconfigLocation, KubeconfigState } from '@/app/models/kubeconfig'
import { fetchKubeconfigLocations } from '@/app/actions/kubeconfig'

const initialState: KubeconfigState = {
  kubeconfigLocations: {
    isLoading: false,
    data: [],
    error: undefined,
  },
  selectedLocation: undefined,
}

export default createReducer(initialState, (builder) => {
  builder
    .addCase(fetchKubeconfigLocations.pending, (state) => ({
      ...state,
      kubeconfigLocations: {
        isLoading: true,
        data: [],
        error: undefined,
      },
    }))
    .addCase(fetchKubeconfigLocations.fulfilled, (state, action: PayloadAction<Array<KubeconfigLocation>>) => ({
      ...state,
      kubeconfigLocations: {
        isLoading: false,
        data: action.payload,
        error: undefined,
      },
    }))
    .addCase(fetchKubeconfigLocations.rejected, (state, action: PayloadAction<unknown>) => ({
      ...state,
      kubeconfigLocations: {
        isLoading: false,
        data: [],
        error: action.payload,
      },
    }))
})
