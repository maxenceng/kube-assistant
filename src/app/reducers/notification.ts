import { createReducer, PayloadAction } from '@reduxjs/toolkit'
import { v4 } from 'uuid'
import { NotificationMessage, NotificationState, NotificationType } from '@/app/models/notification'
import { createAsyncNotification, createNotification, removeNotification } from '@/app/actions/notification'

const initialState: NotificationState = {
  queue: [],
}

export default createReducer(initialState, (builder) => {
  builder
    .addCase(createAsyncNotification.fulfilled, (state: NotificationState, action: PayloadAction<string>) => ({
      ...state,
      queue: [
        ...state.queue,
        {
          uuid: v4(),
          type: NotificationType.SUCCESS,
          message: action.payload,
        },
      ],
    }))
    .addCase(createAsyncNotification.rejected, (state: NotificationState, action: PayloadAction<any>) => ({
      ...state,
      queue: [
        ...state.queue,
        {
          uuid: v4(),
          type: NotificationType.ERROR,
          message: action?.payload ?? 'Unknown error',
        },
      ],
    }))
    .addCase(createNotification, (state: NotificationState, action: PayloadAction<Omit<NotificationMessage, 'uuid'>>) => ({
      ...state,
      queue: [
        ...state.queue,
        {
          ...action.payload,
          uuid: v4(),
        },
      ],
    }))
    .addCase(removeNotification, (state: NotificationState, action: PayloadAction<string>) => ({
      ...state,
      queue: state.queue.filter((notification) => notification.uuid !== action.payload),
    }))
})
