## Agenda

[ ] A Brief Overview of SPIRE and Aegis             » (new/focused speaker deck)                TBD
[ ] Overview the Installed SPIRE Components         » (knss, kgp)                               OK
[ ] Overview the Installed Aegis Components         » (knas, kgp, knd, kgp)                     OK
[ ] What are ClusterSPIFFEIDs                       » (brief definition (from ChatGPT))         OK
[ ] Examine the Current ClusterSPIFFEIDs            » (kid, dsa, dse)                           OK
[ ] Install ClusterSPIFFEIDs We’ll Need             » (iid)                                     OK
[ ] Examine Inspector’s ClusterSPIFFEID             » (kid, din)                                OK
[ ] Install Inspector                               » (iin)                                     OK
[ ] Examine Inspector’s ClusterSPIFFEID             » (din)                                     OK
[ ] Use Case: Workload Using Sidecar                » (ds, dw, isc, twl)                        OK
[ ]   * register a plain secret                     » (rps)                                     OK
[ ]   * register multiple secrets                   » (rms)                                     OK
[ ]   * encrypt a secret                            » (es)                                      OK
[ ]   * register encrypted secret                   » (venv, xes, res)                          OK
[ ] Use Case: Workload Using SDK                    » (ds, dwl, isdk, twl)                      OK
[ ]   * register secret                             » (rps)                                     OK
[ ]   * register multiple secrets                   » (rms)                                     OK
[ ] Use Case: Workload Using Init Container         » (ds, dwo, ic)                             OK
[ ]   * check the workload’s state                  » (kgp -w)                                  OK
[ ]   * register the secret                         » (rks)                                     OK
[ ] Use Case: Transforming Secrets                  » (vsdk, isdk, twl)
[ ]   * register a secret with JSON transformation  » (vjson, rjson)
[ ]   * register a secret with YAML transformation  » (vyaml, ryaml)
[ ]   * remove all secrets                          » (ds)
[ ]   * uninstall the workload                      » (dwo)
[ ] Use Case: Kubernetes Secret Interpolation       » (vinit, iinit)
[ ]   * check the workload’s state                  » (kgp)
[ ]   * register a Secret with K8s Interpolation    » (vk8s, ik8s)
[ ]   * check the workload’s state                  » (kgp)
[ ]   * remove all secrets                          » (ds)
[ ]   * uninstall the workload                      » (dwo)
