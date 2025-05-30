// Code generated by astool. DO NOT EDIT.

package propertyurl

import (
	anyuri "code.superseriousbusiness.org/activity/streams/values/anyURI"
	vocab "code.superseriousbusiness.org/activity/streams/vocab"
	"fmt"
	"net/url"
)

// ActivityStreamsUrlPropertyIterator is an iterator for a property. It is
// permitted to be one of multiple value types. At most, one type of value can
// be present, or none at all. Setting a value will clear the other types of
// values so that only one of the 'Is' methods will return true. It is
// possible to clear all values, so that this property is empty.
type ActivityStreamsUrlPropertyIterator struct {
	xmlschemaAnyURIMember        *url.URL
	activitystreamsLinkMember    vocab.ActivityStreamsLink
	tootHashtagMember            vocab.TootHashtag
	activitystreamsMentionMember vocab.ActivityStreamsMention
	unknown                      interface{}
	alias                        string
	myIdx                        int
	parent                       vocab.ActivityStreamsUrlProperty
}

// NewActivityStreamsUrlPropertyIterator creates a new ActivityStreamsUrl property.
func NewActivityStreamsUrlPropertyIterator() *ActivityStreamsUrlPropertyIterator {
	return &ActivityStreamsUrlPropertyIterator{alias: ""}
}

// deserializeActivityStreamsUrlPropertyIterator creates an iterator from an
// element that has been unmarshalled from a text or binary format.
func deserializeActivityStreamsUrlPropertyIterator(i interface{}, aliasMap map[string]string) (*ActivityStreamsUrlPropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	if m, ok := i.(map[string]interface{}); ok {
		if v, err := mgr.DeserializeLinkActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsUrlPropertyIterator{
				activitystreamsLinkMember: v,
				alias:                     alias,
			}
			return this, nil
		} else if v, err := mgr.DeserializeHashtagToot()(m, aliasMap); err == nil {
			this := &ActivityStreamsUrlPropertyIterator{
				alias:             alias,
				tootHashtagMember: v,
			}
			return this, nil
		} else if v, err := mgr.DeserializeMentionActivityStreams()(m, aliasMap); err == nil {
			this := &ActivityStreamsUrlPropertyIterator{
				activitystreamsMentionMember: v,
				alias:                        alias,
			}
			return this, nil
		}
	}
	if v, err := anyuri.DeserializeAnyURI(i); err == nil {
		this := &ActivityStreamsUrlPropertyIterator{
			alias:                 alias,
			xmlschemaAnyURIMember: v,
		}
		return this, nil
	}
	this := &ActivityStreamsUrlPropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
}

// GetActivityStreamsLink returns the value of this property. When
// IsActivityStreamsLink returns false, GetActivityStreamsLink will return an
// arbitrary value.
func (this ActivityStreamsUrlPropertyIterator) GetActivityStreamsLink() vocab.ActivityStreamsLink {
	return this.activitystreamsLinkMember
}

// GetActivityStreamsMention returns the value of this property. When
// IsActivityStreamsMention returns false, GetActivityStreamsMention will
// return an arbitrary value.
func (this ActivityStreamsUrlPropertyIterator) GetActivityStreamsMention() vocab.ActivityStreamsMention {
	return this.activitystreamsMentionMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this ActivityStreamsUrlPropertyIterator) GetIRI() *url.URL {
	return this.xmlschemaAnyURIMember
}

// GetTootHashtag returns the value of this property. When IsTootHashtag returns
// false, GetTootHashtag will return an arbitrary value.
func (this ActivityStreamsUrlPropertyIterator) GetTootHashtag() vocab.TootHashtag {
	return this.tootHashtagMember
}

// GetType returns the value in this property as a Type. Returns nil if the value
// is not an ActivityStreams type, such as an IRI or another value.
func (this ActivityStreamsUrlPropertyIterator) GetType() vocab.Type {
	if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink()
	}
	if this.IsTootHashtag() {
		return this.GetTootHashtag()
	}
	if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention()
	}

	return nil
}

// GetXMLSchemaAnyURI returns the value of this property. When IsXMLSchemaAnyURI
// returns false, GetXMLSchemaAnyURI will return an arbitrary value.
func (this ActivityStreamsUrlPropertyIterator) GetXMLSchemaAnyURI() *url.URL {
	return this.xmlschemaAnyURIMember
}

// HasAny returns true if any of the different values is set.
func (this ActivityStreamsUrlPropertyIterator) HasAny() bool {
	return this.IsXMLSchemaAnyURI() ||
		this.IsActivityStreamsLink() ||
		this.IsTootHashtag() ||
		this.IsActivityStreamsMention()
}

// IsActivityStreamsLink returns true if this property has a type of "Link". When
// true, use the GetActivityStreamsLink and SetActivityStreamsLink methods to
// access and set this property.
func (this ActivityStreamsUrlPropertyIterator) IsActivityStreamsLink() bool {
	return this.activitystreamsLinkMember != nil
}

// IsActivityStreamsMention returns true if this property has a type of "Mention".
// When true, use the GetActivityStreamsMention and SetActivityStreamsMention
// methods to access and set this property.
func (this ActivityStreamsUrlPropertyIterator) IsActivityStreamsMention() bool {
	return this.activitystreamsMentionMember != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this ActivityStreamsUrlPropertyIterator) IsIRI() bool {
	return this.xmlschemaAnyURIMember != nil
}

// IsTootHashtag returns true if this property has a type of "Hashtag". When true,
// use the GetTootHashtag and SetTootHashtag methods to access and set this
// property.
func (this ActivityStreamsUrlPropertyIterator) IsTootHashtag() bool {
	return this.tootHashtagMember != nil
}

// IsXMLSchemaAnyURI returns true if this property has a type of "anyURI". When
// true, use the GetXMLSchemaAnyURI and SetXMLSchemaAnyURI methods to access
// and set this property.
func (this ActivityStreamsUrlPropertyIterator) IsXMLSchemaAnyURI() bool {
	return this.xmlschemaAnyURIMember != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsUrlPropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	var child map[string]string
	if this.IsActivityStreamsLink() {
		child = this.GetActivityStreamsLink().JSONLDContext()
	} else if this.IsTootHashtag() {
		child = this.GetTootHashtag().JSONLDContext()
	} else if this.IsActivityStreamsMention() {
		child = this.GetActivityStreamsMention().JSONLDContext()
	}
	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this ActivityStreamsUrlPropertyIterator) KindIndex() int {
	if this.IsXMLSchemaAnyURI() {
		return 0
	}
	if this.IsActivityStreamsLink() {
		return 1
	}
	if this.IsTootHashtag() {
		return 2
	}
	if this.IsActivityStreamsMention() {
		return 3
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsUrlPropertyIterator) LessThan(o vocab.ActivityStreamsUrlPropertyIterator) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsXMLSchemaAnyURI() {
		return anyuri.LessAnyURI(this.GetXMLSchemaAnyURI(), o.GetXMLSchemaAnyURI())
	} else if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink().LessThan(o.GetActivityStreamsLink())
	} else if this.IsTootHashtag() {
		return this.GetTootHashtag().LessThan(o.GetTootHashtag())
	} else if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention().LessThan(o.GetActivityStreamsMention())
	}
	return false
}

// Name returns the name of this property: "ActivityStreamsUrl".
func (this ActivityStreamsUrlPropertyIterator) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "ActivityStreamsUrl"
	} else {
		return "ActivityStreamsUrl"
	}
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this ActivityStreamsUrlPropertyIterator) Next() vocab.ActivityStreamsUrlPropertyIterator {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this ActivityStreamsUrlPropertyIterator) Prev() vocab.ActivityStreamsUrlPropertyIterator {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// SetActivityStreamsLink sets the value of this property. Calling
// IsActivityStreamsLink afterwards returns true.
func (this *ActivityStreamsUrlPropertyIterator) SetActivityStreamsLink(v vocab.ActivityStreamsLink) {
	this.clear()
	this.activitystreamsLinkMember = v
}

// SetActivityStreamsMention sets the value of this property. Calling
// IsActivityStreamsMention afterwards returns true.
func (this *ActivityStreamsUrlPropertyIterator) SetActivityStreamsMention(v vocab.ActivityStreamsMention) {
	this.clear()
	this.activitystreamsMentionMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *ActivityStreamsUrlPropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.SetXMLSchemaAnyURI(v)
}

// SetTootHashtag sets the value of this property. Calling IsTootHashtag
// afterwards returns true.
func (this *ActivityStreamsUrlPropertyIterator) SetTootHashtag(v vocab.TootHashtag) {
	this.clear()
	this.tootHashtagMember = v
}

// SetType attempts to set the property for the arbitrary type. Returns an error
// if it is not a valid type to set on this property.
func (this *ActivityStreamsUrlPropertyIterator) SetType(t vocab.Type) error {
	if v, ok := t.(vocab.ActivityStreamsLink); ok {
		this.SetActivityStreamsLink(v)
		return nil
	}
	if v, ok := t.(vocab.TootHashtag); ok {
		this.SetTootHashtag(v)
		return nil
	}
	if v, ok := t.(vocab.ActivityStreamsMention); ok {
		this.SetActivityStreamsMention(v)
		return nil
	}

	return fmt.Errorf("illegal type to set on ActivityStreamsUrl property: %T", t)
}

// SetXMLSchemaAnyURI sets the value of this property. Calling IsXMLSchemaAnyURI
// afterwards returns true.
func (this *ActivityStreamsUrlPropertyIterator) SetXMLSchemaAnyURI(v *url.URL) {
	this.clear()
	this.xmlschemaAnyURIMember = v
}

// clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *ActivityStreamsUrlPropertyIterator) clear() {
	this.xmlschemaAnyURIMember = nil
	this.activitystreamsLinkMember = nil
	this.tootHashtagMember = nil
	this.activitystreamsMentionMember = nil
	this.unknown = nil
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsUrlPropertyIterator) serialize() (interface{}, error) {
	if this.IsXMLSchemaAnyURI() {
		return anyuri.SerializeAnyURI(this.GetXMLSchemaAnyURI())
	} else if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink().Serialize()
	} else if this.IsTootHashtag() {
		return this.GetTootHashtag().Serialize()
	} else if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention().Serialize()
	}
	return this.unknown, nil
}

// ActivityStreamsUrlProperty is the non-functional property "url". It is
// permitted to have one or more values, and of different value types.
type ActivityStreamsUrlProperty struct {
	properties []*ActivityStreamsUrlPropertyIterator
	alias      string
}

// DeserializeUrlProperty creates a "url" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeUrlProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsUrlProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "url"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "url")
	}
	i, ok := m[propName]

	if ok {
		this := &ActivityStreamsUrlProperty{
			alias:      alias,
			properties: []*ActivityStreamsUrlPropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeActivityStreamsUrlPropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeActivityStreamsUrlPropertyIterator(i, aliasMap); err != nil {
				return this, err
			} else if p != nil {
				this.properties = append(this.properties, p)
			}
		}
		// Set up the properties for iteration.
		for idx, ele := range this.properties {
			ele.parent = this
			ele.myIdx = idx
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsUrlProperty creates a new url property.
func NewActivityStreamsUrlProperty() *ActivityStreamsUrlProperty {
	return &ActivityStreamsUrlProperty{alias: ""}
}

// AppendActivityStreamsLink appends a Link value to the back of a list of the
// property "url". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsUrlProperty) AppendActivityStreamsLink(v vocab.ActivityStreamsLink) {
	this.properties = append(this.properties, &ActivityStreamsUrlPropertyIterator{
		activitystreamsLinkMember: v,
		alias:                     this.alias,
		myIdx:                     this.Len(),
		parent:                    this,
	})
}

// AppendActivityStreamsMention appends a Mention value to the back of a list of
// the property "url". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsUrlProperty) AppendActivityStreamsMention(v vocab.ActivityStreamsMention) {
	this.properties = append(this.properties, &ActivityStreamsUrlPropertyIterator{
		activitystreamsMentionMember: v,
		alias:                        this.alias,
		myIdx:                        this.Len(),
		parent:                       this,
	})
}

// AppendIRI appends an IRI value to the back of a list of the property "url"
func (this *ActivityStreamsUrlProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &ActivityStreamsUrlPropertyIterator{
		alias:                 this.alias,
		myIdx:                 this.Len(),
		parent:                this,
		xmlschemaAnyURIMember: v,
	})
}

// AppendTootHashtag appends a Hashtag value to the back of a list of the property
// "url". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsUrlProperty) AppendTootHashtag(v vocab.TootHashtag) {
	this.properties = append(this.properties, &ActivityStreamsUrlPropertyIterator{
		alias:             this.alias,
		myIdx:             this.Len(),
		parent:            this,
		tootHashtagMember: v,
	})
}

// PrependType prepends an arbitrary type value to the front of a list of the
// property "url". Invalidates iterators that are traversing using Prev.
// Returns an error if the type is not a valid one to set for this property.
func (this *ActivityStreamsUrlProperty) AppendType(t vocab.Type) error {
	n := &ActivityStreamsUrlPropertyIterator{
		alias:  this.alias,
		myIdx:  this.Len(),
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	this.properties = append(this.properties, n)
	return nil
}

// AppendXMLSchemaAnyURI appends a anyURI value to the back of a list of the
// property "url". Invalidates iterators that are traversing using Prev.
func (this *ActivityStreamsUrlProperty) AppendXMLSchemaAnyURI(v *url.URL) {
	this.properties = append(this.properties, &ActivityStreamsUrlPropertyIterator{
		alias:                 this.alias,
		myIdx:                 this.Len(),
		parent:                this,
		xmlschemaAnyURIMember: v,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this ActivityStreamsUrlProperty) At(index int) vocab.ActivityStreamsUrlPropertyIterator {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this ActivityStreamsUrlProperty) Begin() vocab.ActivityStreamsUrlPropertyIterator {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this ActivityStreamsUrlProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this ActivityStreamsUrlProperty) End() vocab.ActivityStreamsUrlPropertyIterator {
	return nil
}

// InsertActivityStreamsLink inserts a Link value at the specified index for a
// property "url". Existing elements at that index and higher are shifted back
// once. Invalidates all iterators.
func (this *ActivityStreamsUrlProperty) InsertActivityStreamsLink(idx int, v vocab.ActivityStreamsLink) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsUrlPropertyIterator{
		activitystreamsLinkMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// InsertActivityStreamsMention inserts a Mention value at the specified index for
// a property "url". Existing elements at that index and higher are shifted
// back once. Invalidates all iterators.
func (this *ActivityStreamsUrlProperty) InsertActivityStreamsMention(idx int, v vocab.ActivityStreamsMention) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsUrlPropertyIterator{
		activitystreamsMentionMember: v,
		alias:                        this.alias,
		myIdx:                        idx,
		parent:                       this,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Insert inserts an IRI value at the specified index for a property "url".
// Existing elements at that index and higher are shifted back once.
// Invalidates all iterators.
func (this *ActivityStreamsUrlProperty) InsertIRI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsUrlPropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// InsertTootHashtag inserts a Hashtag value at the specified index for a property
// "url". Existing elements at that index and higher are shifted back once.
// Invalidates all iterators.
func (this *ActivityStreamsUrlProperty) InsertTootHashtag(idx int, v vocab.TootHashtag) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsUrlPropertyIterator{
		alias:             this.alias,
		myIdx:             idx,
		parent:            this,
		tootHashtagMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependType prepends an arbitrary type value to the front of a list of the
// property "url". Invalidates all iterators. Returns an error if the type is
// not a valid one to set for this property.
func (this *ActivityStreamsUrlProperty) InsertType(idx int, t vocab.Type) error {
	n := &ActivityStreamsUrlPropertyIterator{
		alias:  this.alias,
		myIdx:  idx,
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = n
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
	return nil
}

// InsertXMLSchemaAnyURI inserts a anyURI value at the specified index for a
// property "url". Existing elements at that index and higher are shifted back
// once. Invalidates all iterators.
func (this *ActivityStreamsUrlProperty) InsertXMLSchemaAnyURI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ActivityStreamsUrlPropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsUrlProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	for _, elem := range this.properties {
		child := elem.JSONLDContext()
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		for k, v := range child {
			m[k] = v
		}
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API method specifically needed only for alternate implementations
// for go-fed. Applications should not use this method. Panics if the index is
// out of bounds.
func (this ActivityStreamsUrlProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "url" property.
func (this ActivityStreamsUrlProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this ActivityStreamsUrlProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].GetXMLSchemaAnyURI()
			rhs := this.properties[j].GetXMLSchemaAnyURI()
			return anyuri.LessAnyURI(lhs, rhs)
		} else if idx1 == 1 {
			lhs := this.properties[i].GetActivityStreamsLink()
			rhs := this.properties[j].GetActivityStreamsLink()
			return lhs.LessThan(rhs)
		} else if idx1 == 2 {
			lhs := this.properties[i].GetTootHashtag()
			rhs := this.properties[j].GetTootHashtag()
			return lhs.LessThan(rhs)
		} else if idx1 == 3 {
			lhs := this.properties[i].GetActivityStreamsMention()
			rhs := this.properties[j].GetActivityStreamsMention()
			return lhs.LessThan(rhs)
		} else if idx1 == -2 {
			lhs := this.properties[i].GetIRI()
			rhs := this.properties[j].GetIRI()
			return lhs.String() < rhs.String()
		}
	}
	return false
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsUrlProperty) LessThan(o vocab.ActivityStreamsUrlProperty) bool {
	l1 := this.Len()
	l2 := o.Len()
	l := l1
	if l2 < l1 {
		l = l2
	}
	for i := 0; i < l; i++ {
		if this.properties[i].LessThan(o.At(i)) {
			return true
		} else if o.At(i).LessThan(this.properties[i]) {
			return false
		}
	}
	return l1 < l2
}

// Name returns the name of this property ("url") with any alias.
func (this ActivityStreamsUrlProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "url"
	} else {
		return "url"
	}
}

// PrependActivityStreamsLink prepends a Link value to the front of a list of the
// property "url". Invalidates all iterators.
func (this *ActivityStreamsUrlProperty) PrependActivityStreamsLink(v vocab.ActivityStreamsLink) {
	this.properties = append([]*ActivityStreamsUrlPropertyIterator{{
		activitystreamsLinkMember: v,
		alias:                     this.alias,
		myIdx:                     0,
		parent:                    this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependActivityStreamsMention prepends a Mention value to the front of a list
// of the property "url". Invalidates all iterators.
func (this *ActivityStreamsUrlProperty) PrependActivityStreamsMention(v vocab.ActivityStreamsMention) {
	this.properties = append([]*ActivityStreamsUrlPropertyIterator{{
		activitystreamsMentionMember: v,
		alias:                        this.alias,
		myIdx:                        0,
		parent:                       this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependIRI prepends an IRI value to the front of a list of the property "url".
func (this *ActivityStreamsUrlProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*ActivityStreamsUrlPropertyIterator{{
		alias:                 this.alias,
		myIdx:                 0,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependTootHashtag prepends a Hashtag value to the front of a list of the
// property "url". Invalidates all iterators.
func (this *ActivityStreamsUrlProperty) PrependTootHashtag(v vocab.TootHashtag) {
	this.properties = append([]*ActivityStreamsUrlPropertyIterator{{
		alias:             this.alias,
		myIdx:             0,
		parent:            this,
		tootHashtagMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependType prepends an arbitrary type value to the front of a list of the
// property "url". Invalidates all iterators. Returns an error if the type is
// not a valid one to set for this property.
func (this *ActivityStreamsUrlProperty) PrependType(t vocab.Type) error {
	n := &ActivityStreamsUrlPropertyIterator{
		alias:  this.alias,
		myIdx:  0,
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	this.properties = append([]*ActivityStreamsUrlPropertyIterator{n}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
	return nil
}

// PrependXMLSchemaAnyURI prepends a anyURI value to the front of a list of the
// property "url". Invalidates all iterators.
func (this *ActivityStreamsUrlProperty) PrependXMLSchemaAnyURI(v *url.URL) {
	this.properties = append([]*ActivityStreamsUrlPropertyIterator{{
		alias:                 this.alias,
		myIdx:                 0,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "url", regardless of its type. Panics if the index is out of bounds.
// Invalidates all iterators.
func (this *ActivityStreamsUrlProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &ActivityStreamsUrlPropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsUrlProperty) Serialize() (interface{}, error) {
	s := make([]interface{}, 0, len(this.properties))
	for _, iterator := range this.properties {
		if b, err := iterator.serialize(); err != nil {
			return s, err
		} else {
			s = append(s, b)
		}
	}
	// Shortcut: if serializing one value, don't return an array -- pretty sure other Fediverse software would choke on a "type" value with array, for example.
	if len(s) == 1 {
		return s[0], nil
	}
	return s, nil
}

// SetActivityStreamsLink sets a Link value to be at the specified index for the
// property "url". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsUrlProperty) SetActivityStreamsLink(idx int, v vocab.ActivityStreamsLink) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsUrlPropertyIterator{
		activitystreamsLinkMember: v,
		alias:                     this.alias,
		myIdx:                     idx,
		parent:                    this,
	}
}

// SetActivityStreamsMention sets a Mention value to be at the specified index for
// the property "url". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsUrlProperty) SetActivityStreamsMention(idx int, v vocab.ActivityStreamsMention) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsUrlPropertyIterator{
		activitystreamsMentionMember: v,
		alias:                        this.alias,
		myIdx:                        idx,
		parent:                       this,
	}
}

// SetIRI sets an IRI value to be at the specified index for the property "url".
// Panics if the index is out of bounds.
func (this *ActivityStreamsUrlProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsUrlPropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
}

// SetTootHashtag sets a Hashtag value to be at the specified index for the
// property "url". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsUrlProperty) SetTootHashtag(idx int, v vocab.TootHashtag) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsUrlPropertyIterator{
		alias:             this.alias,
		myIdx:             idx,
		parent:            this,
		tootHashtagMember: v,
	}
}

// SetType sets an arbitrary type value to the specified index of the property
// "url". Invalidates all iterators. Returns an error if the type is not a
// valid one to set for this property. Panics if the index is out of bounds.
func (this *ActivityStreamsUrlProperty) SetType(idx int, t vocab.Type) error {
	n := &ActivityStreamsUrlPropertyIterator{
		alias:  this.alias,
		myIdx:  idx,
		parent: this,
	}
	if err := n.SetType(t); err != nil {
		return err
	}
	(this.properties)[idx] = n
	return nil
}

// SetXMLSchemaAnyURI sets a anyURI value to be at the specified index for the
// property "url". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ActivityStreamsUrlProperty) SetXMLSchemaAnyURI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ActivityStreamsUrlPropertyIterator{
		alias:                 this.alias,
		myIdx:                 idx,
		parent:                this,
		xmlschemaAnyURIMember: v,
	}
}

// Swap swaps the location of values at two indices for the "url" property.
func (this ActivityStreamsUrlProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}
