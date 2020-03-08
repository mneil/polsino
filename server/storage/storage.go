package storage

type (
  // IFace stores the shape of storage backends
  IFace interface{
    save() func()
    read() func()
    delete() func()
  }
)
