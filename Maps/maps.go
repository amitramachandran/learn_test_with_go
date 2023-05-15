package maps

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const ErrKeyNotFound = DictionaryErr("cant find the item you searched in dictionary")
const ErrWordDoesNotExists = DictionaryErr("The word doesnt exist in your dictionary")
const ErrKeyAlreadyExists = DictionaryErr("the word you trying to add already exists")

func (d Dictionary) Search(w string) (string, error) {
	meaning, ok := d[w]
	if !ok {
		return meaning, ErrKeyNotFound
	}
	return meaning, nil
}

func (d Dictionary) AddWord(word, description string) error {
	_, err := d.Search(word)

	switch err {
	case ErrKeyNotFound:
		d[word] = description
	case nil:
		return ErrKeyAlreadyExists
	default:
		return err
	}
	return nil

}

func (d Dictionary) UpdateDescription(word, new_description string) error {

	_, err := d.Search(word)

	switch err {
	case ErrKeyNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = new_description
	default:
		return err
	}
	return nil
}

func (d Dictionary) DeleteWord(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case ErrKeyNotFound:
		return ErrWordDoesNotExists
	default:
		return err
	}
	return nil
}
