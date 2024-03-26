import { createAsyncThunk } from '@reduxjs/toolkit'
import axios from '@/app/configs/axios'
import { KubeconfigLocation } from '@/app/models/kubeconfig'

export const fetchKubeconfigLocations = createAsyncThunk<Array<KubeconfigLocation>>(
  'kubeconfig/fetch',
  async () => {
    const response = await axios.get<Array<KubeconfigLocation>>('kubeconfig')
    return response.data
  },
)
