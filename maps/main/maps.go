package main

const (
	ErrKeyNotFound          = ErrorMap("key not found")
	ErrKeyAlreadyRegistered = ErrorMap("key already registered")
	ErrKeyUpdate            = ErrorMap("error updating key value not found")
	ErrKeyDelete            = ErrorMap("error deleting value key not found")
)

type MyDictionary map[string]string
type ErrorMap string

// Search busca um valor associado a uma chave no dicionário
func (md MyDictionary) Search(key string) (string, error) {
	value, exists := md[key]
	if !exists {
		return "", ErrKeyNotFound
	}
	return value, nil
}

// Add adiciona uma nova chave e valor ao dicionário
func (md MyDictionary) Add(key string, value string) error {
	_, err := md.Search(key)
	if err == ErrKeyNotFound {
		md[key] = value
		return nil
	} else if err == nil {
		return ErrKeyAlreadyRegistered
	}
	return err
}

// Update atualiza o valor de uma chave existente no dicionário
func (md MyDictionary) Update(key, value string) error {
	_, err := md.Search(key)
	if err == nil {
		md[key] = value
		return nil
	}
	return ErrKeyUpdate
}

// Delete remove uma chave e seu valor associado do dicionário
func (md MyDictionary) Delete(key string) error {
	_, err := md.Search(key)
	if err == nil {
		delete(md, key)
		return nil
	}
	return ErrKeyDelete
}

// Error implementa a interface error para ErrorMap
func (e ErrorMap) Error() string {
	return string(e)
}
