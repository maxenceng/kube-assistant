export enum NotificationType {
  SUCCESS = 'success',
  WARNING = 'warning',
  ERROR = 'error',
}

export interface AsyncNotification {
  promise: Promise<any>,
  success: string,
  error: string,
}

export interface NotificationMessage {
  uuid: string
  type: NotificationType
  message: string
}

export interface NotificationState {
  queue: Array<NotificationMessage>
}
