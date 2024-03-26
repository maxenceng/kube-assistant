import { RootState } from '@/app/store'

export const getNotifications = (state: RootState) => state.notification.queue
