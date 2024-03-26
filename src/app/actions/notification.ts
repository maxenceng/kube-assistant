import { createAction, createAsyncThunk } from '@reduxjs/toolkit'
import { AsyncNotification, NotificationMessage } from '@/app/models/notification'

export const createAsyncNotification = createAsyncThunk<string, AsyncNotification>(
  'notification/async/add',
  async ({ promise, success, error }: AsyncNotification, { rejectWithValue }) => {
    try {
      await promise
      return success
    } catch (e: unknown) {
      console.error(e)
      return rejectWithValue(error)
    }
  },
)

export const createNotification = createAction<Omit<NotificationMessage, 'uuid'>>('notification/add')

export const removeNotification = createAction<string>('notification/remove')
