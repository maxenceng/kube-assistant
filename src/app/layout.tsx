import { PropsWithChildren } from 'react'
import { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Kube Assistant',
  description: 'Kube Assistant',
}

export default ({ children }: PropsWithChildren) => (
  <html lang="en">
    <body>
      <main>{children}</main>
    </body>
  </html>
)
