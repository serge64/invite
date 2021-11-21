package usecase

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"html/template"
// 	"io"

// 	"github.com/serge64/invite/internal/entity"
// 	"github.com/serge64/invite/internal/repository"
// )

// var (
// 	ErrTokenNotValid  = errors.New("token not valid")
// 	ErrTokenNotFound  = errors.New("token not found")
// 	ErrTokenNotUnique = errors.New("token not unique")
// 	ErrStatusNotValid = errors.New("status not valid")
// 	ErrEmptyDB        = errors.New("empty db")
// )

// type UseCase struct {
// 	context    context.Context
// 	repository repository.Repository
// 	template   *template.Template
// }

// func New(ctx context.Context, r repository.Repository) UseCase {
// 	return UseCase{
// 		context:    ctx,
// 		repository: r,
// 		template:   template.Must(template.ParseGlob("web/template/*.gohtml")),
// 	}
// }

// func (r UseCase) IndexPage(w io.Writer, token string) {
// 	if g, ok := r.repository.Guest.Find(r.context, token); ok {
// 		_ = r.template.ExecuteTemplate(w, "index.gohtml", g)
// 		return
// 	}
// 	_ = r.template.ExecuteTemplate(w, "404.gohtml", nil)
// }

// func (uc UseCase) ChangeStatus(token string, status string) error {
// 	t := entity.Token(token)

// 	if !t.IsValid() {
// 		return ErrTokenNotValid
// 	}

// 	if !entity.StatusValid(status) {
// 		return ErrStatusNotValid
// 	}

// 	g, ok := uc.repository.Guest.Find(uc.context, token)
// 	if !ok {
// 		return ErrTokenNotFound
// 	}

// 	g.Status = entity.ConvertToStatus(status)

// 	err := uc.repository.Guest.Delete(uc.context, token)
// 	if err != nil {
// 		return fmt.Errorf("%w: %s", ErrNoFoundToken, err)
// 	}

// 	_, err = uc.guest.Create(ctx, g)
// 	if err != nil {
// 		return fmt.Errorf("%w: %s", ErrNoUniqueToken, err)
// 	}
// 	return nil
// }

// func (uc UseCase) ChangeAdditionalPerson(t entity.Token, name string) error {
// 	if !t.IsValid() {
// 		return ErrNoValidToken
// 	}

// 	ctx := context.TODO()

// 	g, err := uc.guest.Find(ctx, t)
// 	if err != nil {
// 		return fmt.Errorf("%w: %s", ErrNoFoundToken, err)
// 	}

// 	g.Name2 = name

// 	err = uc.guest.Delete(ctx, g.Token)
// 	if err != nil {
// 		return fmt.Errorf("%w: %s", ErrNoFoundToken, err)
// 	}

// 	_, err = uc.guest.Create(ctx, g)
// 	if err != nil {
// 		return fmt.Errorf("%w: %s", ErrNoUniqueToken, err)
// 	}
// 	return err
// }

// func (uc UseCase) ChangeChoice(t entity.Token, choice1 string, choice2 string) error {
// 	if !t.IsValid() {
// 		return ErrNoValidToken
// 	}

// 	ctx := context.TODO()

// 	g, err := uc.guest.Find(ctx, t)
// 	if err != nil {
// 		return fmt.Errorf("%w: %s", ErrNoFoundToken, err)
// 	}

// 	g.Choice1 = choice1
// 	g.Choice2 = choice2

// 	err = uc.guest.Delete(ctx, g.Token)
// 	if err != nil {
// 		return fmt.Errorf("%w: %s", ErrNoFoundToken, err)
// 	}

// 	_, err = uc.guest.Create(ctx, g)
// 	if err != nil {
// 		return fmt.Errorf("%w: %s", ErrNoUniqueToken, err)
// 	}
// 	return err
// }
