package passport

import (
	"testing"
)

func TestPassport_Validate(t *testing.T) {
	type fields struct {
		Byr string
		Iyr string
		Eyr string
		Hgt string
		Hcl string
		Ecl string
		Pid string
		Cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "exampleValid",
			fields: fields{
				Byr: "1937",
				Iyr: "2017",
				Eyr: "2020",
				Hgt: "183cm",
				Hcl: "#fffffd",
				Ecl: "gry",
				Pid: "860033327",
				Cid: "147"},
			want: true,
		},
		{
			name: "exampleInvalid",
			fields: fields{
				Byr: "1929",
				Iyr: "2013",
				Eyr: "2023",
				Hgt: "",
				Hcl: "#cfa07d",
				Ecl: "amb",
				Pid: "028048884",
				Cid: "350"},
			want: false,
		},
		{
			name: "exampleInteresting",
			fields: fields{
				Byr: "1931",
				Iyr: "2013",
				Eyr: "2024",
				Hgt: "179cm",
				Hcl: "#ae17e1",
				Ecl: "brn",
				Pid: "760753108",
				Cid: ""},
			want: true,
		},
		{
			name: "exampleInvalid2",
			fields: fields{
				Byr: "",
				Iyr: "2011",
				Eyr: "2025",
				Hgt: "59in",
				Hcl: "#cfa07d ",
				Ecl: "brn",
				Pid: "166559648",
				Cid: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Passport{
				Byr: tt.fields.Byr,
				Iyr: tt.fields.Iyr,
				Eyr: tt.fields.Eyr,
				Hgt: tt.fields.Hgt,
				Hcl: tt.fields.Hcl,
				Ecl: tt.fields.Ecl,
				Pid: tt.fields.Pid,
				Cid: tt.fields.Cid,
			}
			if got := p.Validate(); got != tt.want {
				t.Errorf("Passport.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_validateByr(t *testing.T) {
	type fields struct {
		Byr string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "valid1",
			fields: fields{Byr: "2002"},
			want:   true,
		},
		{
			name:   "invalid1",
			fields: fields{Byr: "2003"},
			want:   false,
		},
		{
			name:   "invalid2",
			fields: fields{Byr: "1919"},
			want:   false,
		},
		{
			name:   "invalid3",
			fields: fields{Byr: "xyz"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Passport{
				Byr: tt.fields.Byr,
			}
			if got := p.validateByr(); got != tt.want {
				t.Errorf("Passport.validateByr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_validateIyr(t *testing.T) {
	type fields struct {
		Iyr string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "valid1",
			fields: fields{Iyr: "2020"},
			want:   true,
		},
		{
			name:   "invalid1",
			fields: fields{Iyr: "2021"},
			want:   false,
		},
		{
			name:   "invalid2",
			fields: fields{Iyr: "2009"},
			want:   false,
		},
		{
			name:   "invalid3",
			fields: fields{Iyr: "asdf"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Passport{
				Iyr: tt.fields.Iyr,
			}
			if got := p.validateIyr(); got != tt.want {
				t.Errorf("Passport.validateIyr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_validateEyr(t *testing.T) {
	type fields struct {
		Eyr string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "valid1",
			fields: fields{Eyr: "2030"},
			want:   true,
		},
		{
			name:   "invalid1",
			fields: fields{Eyr: "2031"},
			want:   false,
		},
		{
			name:   "invalid2",
			fields: fields{Eyr: "2019"},
			want:   false,
		},
		{
			name:   "invalid3",
			fields: fields{Eyr: "asdf"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Passport{
				Eyr: tt.fields.Eyr,
			}
			if got := p.validateEyr(); got != tt.want {
				t.Errorf("Passport.validateEyr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_validateHgt(t *testing.T) {
	type fields struct {
		Hgt string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "valid1",
			fields: fields{Hgt: "60in"},
			want:   true,
		},
		{
			name:   "valid2",
			fields: fields{Hgt: "190cm"},
			want:   true,
		},
		{
			name:   "invalid1",
			fields: fields{Hgt: "190in"},
			want:   false,
		},
		{
			name:   "invalid2",
			fields: fields{Hgt: "190"},
			want:   false,
		},
		{
			name:   "invalid3",
			fields: fields{Hgt: "cm"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Passport{
				Hgt: tt.fields.Hgt,
			}
			if got := p.validateHgt(); got != tt.want {
				t.Errorf("Passport.validateHgt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_validateHcl(t *testing.T) {
	type fields struct {
		Hcl string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "valid1",
			fields: fields{Hcl: "#123abc"},
			want:   true,
		},
		{
			name:   "invalid1",
			fields: fields{Hcl: "#123abz"},
			want:   false,
		},
		{
			name:   "invalid2",
			fields: fields{Hcl: "123abc"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Passport{
				Hcl: tt.fields.Hcl,
			}
			if got := p.validateHcl(); got != tt.want {
				t.Errorf("Passport.validateHcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_validateEcl(t *testing.T) {
	type fields struct {
		Ecl string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "valid1",
			fields: fields{Ecl: "brn"},
			want:   true,
		},
		{
			name:   "invalid1",
			fields: fields{Ecl: "wat"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Passport{
				Ecl: tt.fields.Ecl,
			}
			if got := p.validateEcl(); got != tt.want {
				t.Errorf("Passport.validateEcl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_validatePid(t *testing.T) {
	type fields struct {
		Pid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "valid1",
			fields: fields{Pid: "000000001"},
			want:   true,
		},
		{
			name:   "invalid1",
			fields: fields{Pid: "0123456789"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Passport{
				Pid: tt.fields.Pid,
			}
			if got := p.validatePid(); got != tt.want {
				t.Errorf("Passport.validatePid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassport_ValidateMore(t *testing.T) {
	type fields struct {
		Byr string
		Iyr string
		Eyr string
		Hgt string
		Hcl string
		Ecl string
		Pid string
		Cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "exampleInvalid1",
			fields: fields{
				Byr: "1926",
				Iyr: "2018",
				Eyr: "1972",
				Hgt: "170",
				Hcl: "#18171d",
				Ecl: "amb",
				Pid: "186cm",
				Cid: "100"},
			want: false,
		},
		{
			name: "exampleInvalid2",
			fields: fields{
				Byr: "1946",
				Iyr: "2019",
				Eyr: "1967",
				Hgt: "170",
				Hcl: "#602927",
				Ecl: "grn",
				Pid: "012533040",
				Cid: "100"},
			want: false,
		},
		{
			name: "exampleValid1",
			fields: fields{
				Byr: "1980",
				Iyr: "2012",
				Eyr: "2030",
				Hgt: "74in",
				Hcl: "#623a2f",
				Ecl: "grn",
				Pid: "087499704",
				Cid: ""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Passport{
				Byr: tt.fields.Byr,
				Iyr: tt.fields.Iyr,
				Eyr: tt.fields.Eyr,
				Hgt: tt.fields.Hgt,
				Hcl: tt.fields.Hcl,
				Ecl: tt.fields.Ecl,
				Pid: tt.fields.Pid,
				Cid: tt.fields.Cid,
			}
			if got := p.ValidateMore(); got != tt.want {
				t.Errorf("Passport.ValidateMore() = %v, want %v", got, tt.want)
			}
		})
	}
}
