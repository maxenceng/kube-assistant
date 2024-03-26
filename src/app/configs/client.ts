import k8s, { CoreV1Api } from '@kubernetes/client-node'

export const connect = (kubeconfig: string): CoreV1Api => {
  const kc = new k8s.KubeConfig()
  kc.loadFromString(kubeconfig)

  return kc.makeApiClient(k8s.CoreV1Api)
}
