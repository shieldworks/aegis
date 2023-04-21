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
[ ]   * register a plain secret                     » (rps)
[ ]   * register multiple secrets                   » (rms)
[ ]   * encrypt a secret                            » (es, venv (contains multiple secrets)))
[ ]   * register encrypted secret                   » (ves, res)
[ ]   * encrypt multiple secrets                    » (vems, ems, venv)
[ ]   * register multiple encrypted secrets         » (vrmes, rmes)
[ ]   * remove all secrets                          » (vds, ds)
[ ]   * uninstall the workload                      » (vdwo, dwo)
[ ] Use Case: Workload Using SDK                    » (vsdk, isdk, twl)
[ ]   * register secret                             » (rps)
[ ]   * remove all secrets                          » (ds)
[ ]   * uninstall the workload                      » (dwo)
[ ] Use Case: Workload Using Init Container         » (vinit, iinit, twl)
[ ]   * check the workload’s state                  » (kgp)
[ ]   * register the secret                         » (rps)
[ ]   * check the workload’s state                  » (kgp, twl)
[ ]   * remove all secrets                          » (ds)
[ ]   * uninstall the workload                      » (dwl)
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
