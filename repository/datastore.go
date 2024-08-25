package repository

import (
	"errors"
	"slices"

	"github.com/john-k-antony/richpanel-blogposts-golang/models/apimodels"
)

//var dataStore = make(map[string]*apimodels.Blogpost)

var dataStore []*apimodels.Blogpost

func Put(data *apimodels.Blogpost) {
	if data != nil {
		//dataStore[*data.ID] = data
		dataStore = append(dataStore, data)
	}
}

func Get(id string) (*apimodels.Blogpost, error) {
	//blogpost, ok := dataStore[id]
	// blogpost, ok := dataStore.Load(id)
	// if ok {
	// 	return blogpost, nil
	// } else {
	// 	return nil, errors.New("Not Found")
	// }

	index := slices.IndexFunc(dataStore, func(x *apimodels.Blogpost) bool {
		if x != nil && *x.ID == id {
			return true
		} else {
			return false
		}
	})

	if index > -1 {
		return dataStore[index], nil
	} else {
		return nil, errors.New("Not Found")
	}
}

func Remove(id string) (*apimodels.Blogpost, error) {
	data, err := Get(id)

	if err != nil {
		return nil, err
	}

	//delete(dataStore, id)
	dataStore = slices.DeleteFunc(dataStore, func(x *apimodels.Blogpost) bool {
		if x != nil && *x.ID == id {
			return true
		} else {
			return false
		}
	})

	return data, nil
}

func List(offset int, limit int) []*apimodels.Blogpost {
	// vals := slices.Collect(maps.Values(dataStore))
	vals := dataStore
	start := 0
	if offset > 0 {
		start = 0
	}

	maxRecordSize := 100
	updatedLimit := maxRecordSize

	if limit > 0 {
		updatedLimit = limit
	}
	if updatedLimit > maxRecordSize {
		updatedLimit = maxRecordSize
	}

	end := start + updatedLimit

	totalLen := len(vals)

	if end > totalLen {
		end = totalLen
	}

	return vals[start:end]
}

func Size() int {
	return len(dataStore)
}
